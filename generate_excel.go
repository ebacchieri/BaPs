//go:build generate_excel
// +build generate_excel

package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	pb "github.com/gucooing/BaPs/common/server_only"

	"google.golang.org/protobuf/proto"
)

func main() {
	excel := &pb.Excel{}
	folders := []string{"./resources/Excel", "./resources/ExcelDB"}

	// Collect all JSON files grouped by base field name with order
	type fileInfo struct {
		path  string
		order int
	}
	fileGroups := make(map[string][]fileInfo)

	for _, folder := range folders {
		err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
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
			fmt.Fprintf(os.Stderr, "Error walking folder %s: %v\n", folder, err)
			os.Exit(1)
		}
	}

	v := reflect.ValueOf(excel).Elem()

	for baseName, files := range fileGroups {
		field := v.FieldByName(baseName)
		if !field.IsValid() {
			fmt.Printf("Skipping %s: not a valid field in Excel\n", baseName)
			continue
		}

		sliceType := field.Type()
		if sliceType.Kind() != reflect.Slice {
			fmt.Printf("Skipping %s: not a valid slice field in Excel\n", baseName)
			continue
		}

		elemType := sliceType.Elem()
		if elemType.Kind() == reflect.Ptr {
			elemType = elemType.Elem()
		}
		if elemType.Kind() != reflect.Struct {
			fmt.Printf("Skipping %s: slice element is not a struct (type: %s)\n", baseName, elemType.String())
			continue
		}

		// Sort files by order ascending
		sort.Slice(files, func(i, j int) bool {
			return files[i].order < files[j].order
		})

		// Accumulate merged items from all files
		var mergedItems []map[string]interface{}

		for _, fi := range files {
			raw, err := ioutil.ReadFile(fi.path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Read %s failed: %v\n", fi.path, err)
				os.Exit(1)
			}

			// Replace None_ → None
			fixedJSON := strings.ReplaceAll(string(raw), `"None_"`, `"None"`)

			// Try to unmarshal as array first
			var items []map[string]interface{}
			if err := json.Unmarshal([]byte(fixedJSON), &items); err != nil {
				// Try as object with DataList wrapper
				var wrapper struct {
					DataList []map[string]interface{} `json:"DataList"`
				}
				if err2 := json.Unmarshal([]byte(fixedJSON), &wrapper); err2 != nil {
					fmt.Fprintf(os.Stderr, "JSON unmarshal %s failed: %v\n", fi.path, err)
					os.Exit(1)
				}
				items = wrapper.DataList
			}

			// Fix bool fields for this file's items
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

		// Marshal merged items and unmarshal into actual slice type
		cleanJSON, err := json.Marshal(mergedItems)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Re-marshal failed: %v\n", err)
			os.Exit(1)
		}

		slicePtr := reflect.New(sliceType).Interface()
		if err := json.Unmarshal(cleanJSON, slicePtr); err != nil {
			fmt.Fprintf(os.Stderr, "Final unmarshal %s failed: %v\n", baseName, err)
			os.Exit(1)
		}

		field.Set(reflect.ValueOf(slicePtr).Elem())
		fmt.Printf("Loaded %s (merged %d files)\n", baseName, len(files))
	}

	// Serialize full Excel struct to protobuf binary
	outPath := "./data/Excel.bin"
	binData, err := proto.Marshal(excel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal proto: %v\n", err)
		os.Exit(1)
	}
	if err := ioutil.WriteFile(outPath, binData, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write bin file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Excel.bin generated successfully.")
}

// parseBaseNameAndOrder extracts base field name and order from filename
// E.g. "AcademyMessanger1ExcelTable" => ("AcademyMessangerExcelTable", 1)
func parseBaseNameAndOrder(name string) (baseName string, order int) {
	const suffix = "ExcelTable"
	if !strings.HasSuffix(name, suffix) {
		return name, 0
	}

	base := strings.TrimSuffix(name, suffix)

	// Extract trailing digits from base
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

// isBoolField reports whether a JSON key corresponds to a bool field in the given struct type.
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
