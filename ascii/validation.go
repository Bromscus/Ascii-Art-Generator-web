package ascii

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func IsValid(s string) (string, bool) {
	array := []rune(s)
	var sr string
	if len(s) < 1 {
		sr = "Enter an input"
		return sr, false
	}

	if len(os.Args) == 2 && os.Args[1] == "" {
		return "Enter an input", false
	}

	if len(array) > 100 {
		sr = "You have exceeded the word limit of 100 character limit. Try again with less characters."
		return sr, false
	}

	for _, v := range array {
		if v < 32 || v > 126 {
			if v == '\n' || v == '\t' || v == '\r' {
				continue
			} else {
				sr = "Invalid character:" + string(v)
				return sr, false
			}
		}
	}

	return sr, true
}

func CheckFile(fileName string, expectedHash string) bool {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return false
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash) == expectedHash
}
