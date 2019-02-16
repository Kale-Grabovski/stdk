package service

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

type CryptService struct{}

type CryptServiceInterface interface {
	HashFile(filePath string) (string, error)
}

func NewCryptService() CryptServiceInterface {
	return &CryptService{}
}

func (h *CryptService) HashFile(filePath string) (string, error) {
	var returnMD5String string

	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}
