package handbook

import (
	"fmt"
	"github.com/gucooing/BaPs/gdconf"
	"github.com/gucooing/BaPs/pkg/logger"
	"os"
	"path/filepath"
)

func NewHandbook() {
	gen := false
	for _, v := range os.Args {
		if v == "handbook" {
			gen = true
		}
	}
	if !gen {
		return
	}
	logger.Info("开始生成handbook.txt")

	file := "handbook\n"

	file += "\nMemoryLobby\n"
	for _, v := range gdconf.GetMemoryLobbyExcelList() {
		file += fmt.Sprintf("%v : %s\n", v.Id, v.PrefabName)
	}

	file += "\nSticker\n"
	for _, v := range gdconf.GetStickerPageContentExcelList() {
		file += fmt.Sprintf("%v\n", v.Id)
	}

	file += "\nEmblem\n"
	for _, v := range gdconf.GetEmblemExcelList() {
		file += fmt.Sprintf("%v\n", v.Id)
	}

	file += "\nItem\n"
	for k, vs := range gdconf.GetItemExcelCategoryMap() {
		file += "  " + k + "\n"
		for _, v := range vs {
			file += fmt.Sprintf("  %v\n", v.Id)
		}
	}

	file += "\nCharacter\n"
	for _, v := range gdconf.GetCharacterMap() {
		file += fmt.Sprintf("%v : %s\n", v.Id, v.DevName)
	}

	file += "\nEquipment\n"
	for _, v := range gdconf.GetEquipmentExcelMap() {
		file += fmt.Sprintf("%v : %s\n", v.Id, v.EquipmentCategory)
	}

	file += "\nFurniture\n"
	for _, v := range gdconf.GetFurnitureExcelTableMap() {
		file += fmt.Sprintf("%v\n", v.Id)
	}

	file += "\nIdCardBackground:\n"
	for _, v := range gdconf.GetIdCardBackgroundExcelList() {
		file += fmt.Sprintf("%v\n", v.Id)
	}

	if saveFile([]byte(file), "./handbook.txt") != nil {
		logger.Error("生成handbook.txt失败")
	} else {
		logger.Info("生成handbook.txt完成")
	}

}

func saveFile(bin []byte, path string) error {
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), 0777)
	}

	return os.WriteFile(path, bin, 0644)
}
