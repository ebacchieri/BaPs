package mx

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io"
	"math/rand"

	"github.com/go-resty/resty/v2"

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
	if len(bin) <= 14 {
		return nil, errors.New("binary too short")
	}
	r := bytes.NewReader(bin)
	r.Read(make([]byte, 4)) // CRC
	r.Read(make([]byte, 4)) // Type conversion
	keyLen, _ := r.ReadByte()
	ivLen, _ := r.ReadByte()

	useAES := keyLen != 0 && ivLen != 0
	aesKey, aesIv := []byte{}, []byte{}
	if useAES {
		if keyLen > 0 {
			aesKey = make([]byte, keyLen)
			r.Read(aesKey)
		}
		if ivLen > 0 {
			aesIv = make([]byte, ivLen)
			r.Read(aesIv)
		}
		r.Read(make([]byte, 4)) // payload len
	} else {
		r.Read(make([]byte, 4)) // payload len
	}

	headerSize := 14 + len(aesKey) + len(aesIv)

	payload := bin[headerSize:]

	alg.Xor(payload, []byte{0xD9})
	z, err := gzip.NewReader(bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer z.Close()
	plain, err := io.ReadAll(z)
	if err != nil || len(aesKey) != 16 || len(aesIv) != 16 {
		return plain, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	if len(plain)%aes.BlockSize != 0 {
		return nil, errors.New("aes payload not multiple of block size")
	}
	decrypted := make([]byte, len(plain))
	cipher.NewCBCDecrypter(block, aesIv).CryptBlocks(decrypted, plain)
	padLen := int(decrypted[len(decrypted)-1])
	if padLen > 0 && padLen <= aes.BlockSize {
		decrypted = decrypted[:len(decrypted)-padLen]
	}
	return decrypted, nil
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
