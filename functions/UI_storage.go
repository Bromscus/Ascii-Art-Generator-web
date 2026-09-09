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
	for {
		num := 2

		switch num {
		case 0:

			banner = "formats/standard.txt"
			hash = StandardHash
		case 1:

			banner = "formats/shadow.txt"
			hash = ShadowHash

		case 2:

			banner = "formats/thinkertoy.txt"
			hash = ThinkertoyHash
		}
		if !ascii.CheckFile(banner, hash) {
			continue
		}
		break
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
