package gdconf

import (
	"github.com/gucooing/BaPs/pkg/logger"
	"os"
	"strings"
	"testing"
)

var dataPath = "./data"
var resPath = "./resources"

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestLoadExcel(t *testing.T) {
	logger.InitLogger("test_gdconf", strings.ToUpper("debug"))
	g := &GameConfig{
		dataPath: dataPath,
		resPath:  resPath,
	}
	GC = g
	g.LoadExcel()
}

func BenchmarkLoadExcel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.InitLogger("test_gdconf", strings.ToUpper("debug"))
		g := &GameConfig{
			dataPath: dataPath,
			resPath:  resPath,
		}
		GC = g
		g.LoadExcel()
	}
}
