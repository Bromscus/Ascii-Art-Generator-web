package functions

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var (
	ShadowHash     [32]byte
	StandardHash   [32]byte
	ThinkertoyHash [32]byte
)

func LoadHashes() {
	ShadowHash = shadowHash()
	StandardHash = standardHash()
	ThinkertoyHash = thinkertoyHash()
}

func shadowHash() [32]byte {
	data, err := os.ReadFile("formats/shadow.txt")
	if err != nil {
		fmt.Println(err)
		return [32]byte{}
	}

	return sha256.Sum256(data)
}

func standardHash() [32]byte {
	data, err := os.ReadFile("formats/standard.txt")
	if err != nil {
		fmt.Println(err)
		return [32]byte{}
	}

	return sha256.Sum256(data)
}

func thinkertoyHash() [32]byte {
	data, err := os.ReadFile("formats/thinkertoy.txt")
	if err != nil {
		fmt.Println(err)
		return [32]byte{}
	}

	return sha256.Sum256(data)
}
