package mx

import (
	"bytes"
	"compress/gzip"
	"errors"
	"github.com/go-resty/resty/v2"
	"io"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/BaPs/pkg/alg"
)

func GetFormMx(c *gin.Context) ([]byte, error) {
	file, err := c.FormFile("mx")
	if err != nil {
		return nil, err
	}
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}
	bin := make([]byte, file.Size)
	_, err = fileContent.Read(bin)
	if err != nil {
		return nil, err
	}
	// 下面是解密
	return DeMx(bin)
}

func DeMx(bin []byte) ([]byte, error) {
	if len(bin) <= 12 {
		return nil, errors.New("binary too short")
	}
	alg.Xor(bin, []byte{0xD9})
	z, err := gzip.NewReader(bytes.NewReader(bin[12:]))
	if err != nil {
		return nil, err
	}
	defer z.Close()
	p, err := io.ReadAll(z)
	return p, err
}

func SetFormMx(req *resty.Request, bin []byte) error {
	encodedData, err := EncodeMx(bin)
	if err != nil {
		return err
	}
	req.SetFileReader("mx", "data.mx", bytes.NewReader(encodedData))

	return nil
}

func EncodeMx(originalData []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(originalData); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	compressedData := buf.Bytes()
	header := make([]byte, 12)
	if _, err := rand.Read(header); err != nil {
		return nil, err
	}
	preXorData := append(header, compressedData...)
	xorKey := byte(0xD9)
	for i := range preXorData {
		preXorData[i] ^= xorKey
	}

	return preXorData, nil
}
