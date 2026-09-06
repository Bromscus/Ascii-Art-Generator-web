package functions

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var (
	ShadowHash     string
	StandardHash   string
	ThinkertoyHash string
)

func LoadHashes() {
	ShadowHash = shadowHash()
	StandardHash = standardHash()
	ThinkertoyHash = thinkertoyHash()
}

func shadowHash() string {
	data, err := os.ReadFile("formats/shadow.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func standardHash() string {
	data, err := os.ReadFile("formats/standard.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func thinkertoyHash() string {
	data, err := os.ReadFile("formats/thinkertoy.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}
