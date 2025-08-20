//go:build !debug
// +build !debug

package gdconf

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
	pb "google.golang.org/protobuf/proto"
)

var (
	excelPath   = "./resources/Excel"
	excelDbPath = "./resources/ExcelDB"
)

func (g *GameConfig) LoadExcel() {
retry:
	info, err := os.Stat(g.dataPath)
	if err != nil || !info.IsDir() {
		panic("找不到文件夹: " + g.dataPath)
	}
	if !strings.HasSuffix(g.dataPath, "/") {
		g.dataPath += "/"
	}

	binFile := g.dataPath + "Excel.bin"
	data, err := os.ReadFile(binFile)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info("Excel.bin not found; generating from dumped JSONs...")
			if genErr := GenerateExcelBin(excelPath, excelDbPath, binFile); genErr != nil {
				logger.Error("生成 Excel.bin 失败: %v", genErr)
				return
			}
			logger.Info("Excel.bin 生成成功，重新加载...")
			goto retry
		}
		logger.Error("无法读取 Excel.bin: %v", err)
		return
	}

	g.Excel = new(sro.Excel)
	if err := pb.Unmarshal(data, g.Excel); err != nil {
		panic("解析 Excel.bin 失败，请检查版本是否匹配")
	}
}

func GenerateExcelBin(excelPath, excelDbPath, outBinPath string) error {
	excel := &sro.Excel{}
	folders := []string{excelPath, excelDbPath}

	// Collect all JSON files by base field name (without trailing digits)
	type fileInfo struct {
		path  string
		order int // ordering by trailing number suffix
	}
	fileGroups := make(map[string][]fileInfo)

	for _, folder := range folders {
		err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				logger.Error("walk %s failed: %v", path, err)
				return err
			}
			if info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
				return nil
			}

			name := strings.TrimSuffix(info.Name(), ".json")
			baseName, order := parseBaseNameAndOrder(name)
			fileGroups[baseName] = append(fileGroups[baseName], fileInfo{path, order})
			return nil
		})

		if err != nil {
			logger.Error("error walking folder %s: %v", folder, err)
			return err
		}
	}

	v := reflect.ValueOf(excel).Elem()

	for baseName, files := range fileGroups {
		field := v.FieldByName(baseName)
		if !field.IsValid() {
			continue // skip unknown fields
		}

		sliceType := field.Type()
		if sliceType.Kind() != reflect.Slice {
			continue
		}

		elemType := sliceType.Elem()
		if elemType.Kind() == reflect.Ptr {
			elemType = elemType.Elem()
		}
		if elemType.Kind() != reflect.Struct {
			continue
		}

		// Sort files by order ascending
		sort.Slice(files, func(i, j int) bool {
			return files[i].order < files[j].order
		})

		// Accumulate all items from files
		var mergedItems []map[string]interface{}

		for _, fi := range files {
			raw, err := ioutil.ReadFile(fi.path)
			if err != nil {
				logger.Error("read %s failed: %v", fi.path, err)
				return err
			}

			fixedJSON := strings.ReplaceAll(string(raw), `"None_"`, `"None"`)

			var items []map[string]interface{}

			// Try unmarshal as array first
			err = json.Unmarshal([]byte(fixedJSON), &items)
			if err != nil {
				// Try as object with DataList
				var wrapper struct {
					DataList []map[string]interface{} `json:"DataList"`
				}
				err2 := json.Unmarshal([]byte(fixedJSON), &wrapper)
				if err2 != nil {
					logger.Error("json unmarshal %s failed: %v", fi.path, err)
					return err
				}
				items = wrapper.DataList
			}

			if len(items) == 0 {
				continue // skip empty
			}

			// Fix bool fields inside items
			for _, item := range items {
				for key, val := range item {
					if isBoolField(elemType, key) {
						switch v := val.(type) {
						case float64:
							item[key] = v != 0
						case int:
							item[key] = v != 0
						}
					}
				}
			}

			mergedItems = append(mergedItems, items...)
		}

		// Marshal mergedItems into JSON and unmarshal into proto field slice
		cleanJSON, err := json.Marshal(mergedItems)
		if err != nil {
			logger.Error("re-marshal failed: %v", err)
			return err
		}

		slicePtr := reflect.New(sliceType).Interface()
		if err := json.Unmarshal(cleanJSON, slicePtr); err != nil {
			logger.Error("final unmarshal %s failed: %v", baseName, err)
			return err
		}

		field.Set(reflect.ValueOf(slicePtr).Elem())
		logger.Info("Loaded %s (merged %d files)", baseName, len(files))
	}

	binData, err := pb.Marshal(excel)
	if err != nil {
		logger.Error("failed to marshal proto: %v", err)
		return err
	}
	if err := ioutil.WriteFile(outBinPath, binData, 0644); err != nil {
		logger.Error("failed to write bin file: %v", err)
		return err
	}
	return nil
}

// parseBaseNameAndOrder extracts the base field name and optional trailing number order.
// e.g. "AcademyMessangerExcelTable" => ("AcademyMessangerExcelTable", 0)
//
//	"AcademyMessanger1ExcelTable" => ("AcademyMessangerExcelTable", 1)
//	"AcademyMessanger12ExcelTable" => ("AcademyMessangerExcelTable", 12)
func parseBaseNameAndOrder(name string) (baseName string, order int) {
	// Find trailing digits before "ExcelTable"
	const suffix = "ExcelTable"
	if !strings.HasSuffix(name, suffix) {
		return name, 0
	}

	base := strings.TrimSuffix(name, suffix)
	// Extract trailing digits at end of base (if any)
	i := len(base) - 1
	for i >= 0 && base[i] >= '0' && base[i] <= '9' {
		i--
	}

	numPart := base[i+1:]
	baseName = base[:i+1] + suffix

	if numPart != "" {
		n, err := strconv.Atoi(numPart)
		if err == nil {
			order = n
			return
		}
	}

	order = 0
	return
}

func isBoolField(t reflect.Type, jsonKey string) bool {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		name := strings.Split(tag, ",")[0]
		if name == "" {
			name = f.Name
		}
		if strings.EqualFold(jsonKey, name) && f.Type.Kind() == reflect.Bool {
			return true
		}
	}
	return false
}

func loadExcelFile[T any](path string, table *[]*T) {
	*table = make([]*T, 0)
}
