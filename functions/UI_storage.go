package functions

import (
	"fmt"
	"os"

	"web/ascii"
)

func Header(text string) string {
	banners := []struct {
		file string
		hash [32]byte
	}{
		{"formats/thinkertoy.txt", ThinkertoyHash},
		{"formats/standard.txt", StandardHash},
		{"formats/shadow.txt", ShadowHash},
	}

	for _, candidate := range banners {
		if !ascii.CheckFile(candidate.file, candidate.hash) {
			continue
		}

		data, err := os.ReadFile(candidate.file)
		if err != nil {
			fmt.Println(err)
			continue
		}

		return ascii.Process(text, data)
	}

	return ""
}
