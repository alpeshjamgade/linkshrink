package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"hash/crc32"
)

func GetUUID() string {
	return uuid.New().String()
}

func GenerateHash(url string) string {
	dataBytes := []byte(url)
	encoded := base64.StdEncoding.EncodeToString(dataBytes)
	return encoded
}

func GenerateCRC32Hash(data string) string {
	dataBytes := []byte(data)

	hash := crc32.NewIEEE()
	hash.Write(dataBytes)

	checksum := hash.Sum32()
	return fmt.Sprintf("%08x", checksum)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
