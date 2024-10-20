package cores

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func EncodeHexToString(data []byte) string {
	return hex.EncodeToString(data)
}

func DecodeHexToBytes(data string) ([]byte, error) {
	return hex.DecodeString(data)
}

func GetHashSha256(data []byte) []byte {
	hash := sha256.New()
	buff := hash.Sum(data)
	return buff
}

type StrOrBuffImpl interface {
	~[]byte | ~string
}

func BytesEquals[V StrOrBuffImpl](data V, buff V) bool {
	size := len(data)
	if size != len(buff) {
		return false
	}
	for i := 0; i < size; i++ {
		if data[i] != buff[i] {
			return false
		}
	}
	return true
}

func StringEquals(data string, buff string) bool {
	return BytesEquals(data, buff)
}

func HashSha256Compare(data []byte, hash []byte) bool {
	temp := GetHashSha256(data)
	return BytesEquals(temp, hash)
}

func GetHashSha256ToString(data []byte) string {
	temp := GetHashSha256(data)
	return EncodeHexToString(temp)
}

func HashSha256StringCompare(data []byte, hash string) bool {
	temp := GetHashSha256ToString(data)
	return StringEquals(temp, hash)
}

func GetHashSha512(data []byte) []byte {
	hash := sha512.New()
	buff := hash.Sum(data)
	return buff
}

func HashSha512Compare(data []byte, hash []byte) bool {
	temp := GetHashSha512(data)
	return BytesEquals(temp, hash)
}

func GetHashSha512ToString(data []byte) string {
	temp := GetHashSha512(data)
	return EncodeHexToString(temp)
}

func HashSha512StringCompare(data []byte, hash string) bool {
	temp := GetHashSha512ToString(data)
	return StringEquals(temp, hash)
}

func ViperJwtConfigUnmarshal(key string) (*JwtConfig, error) {
	var err error
	KeepVoid(err)

	jwtConfig := NewJwtConfig()
	if err = viper.UnmarshalKey(key, jwtConfig); err != nil {
		return nil, err
	}

	return jwtConfig, nil
}

func EnsureDirAndFile(filePath string) error {
	var err error
	var fileInfo os.FileInfo
	var file *os.File
	KeepVoid(err, fileInfo, file)

	pathDir := filepath.Dir(filePath)
	pathFile := filepath.Base(filePath)

	// Check if the directory exists, and create it if it doesn't
	if fileInfo, err = os.Stat(pathDir); os.IsNotExist(err) {
		if err = os.MkdirAll(pathDir, os.ModePerm); err != nil {
			return NewThrow(fmt.Sprintf("failed to create directory: %s", pathDir), err)
		}
		fmt.Printf("Directory %s created\n", pathDir)
	} else {
		fmt.Printf("Directory %s already exists\n", pathDir)
	}

	// Check if the file exists, and create it if it doesn't
	if fileInfo, err = os.Stat(filePath); os.IsNotExist(err) {
		if file, err = os.Create(filePath); err != nil {
			return NewThrow(fmt.Sprintf("failed to create file: %s", pathFile), err)
		}
		NoErr(file.Close())
		fmt.Printf("File %s created\n", pathFile)
	} else {
		fmt.Printf("File %s already exists\n", pathFile)
	}

	return nil
}
