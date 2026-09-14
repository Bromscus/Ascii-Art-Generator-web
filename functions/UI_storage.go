package functions

import (
	"fmt"
	"os"

	"web/ascii"
)

var (
	hash   [32]byte
	banner string
)

func Header() string {
	banners := []struct {
		file string
		hash [32]byte
	}{
		{"formats/thinkertoy.txt", ThinkertoyHash},
		{"formats/standard.txt", StandardHash},
		{"formats/shadow.txt", ShadowHash},
	}

	banner = ""
	for _, candidate := range banners {
		if ascii.CheckFile(candidate.file, candidate.hash) {
			banner = candidate.file
			hash = candidate.hash
			break
		}
	}

	if banner == "" {
		return ""
	}

	title := ""
	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	title = ascii.Process("Ascii Art Generator", data)
	return title
}
