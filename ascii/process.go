package ascii

import (
	"strings"
)

var num rune

func Process(text string, data []byte) string {
	text = strings.ReplaceAll(text, "\t", "    ")
	table := [8][]string{}
	sentence := []rune(text)

	result := ""

	num = 32
	counter := 1

	lines := strings.Split(string(data), "\n")
	flag := true

	for i := 0; i < len(sentence); i++ {
		v := sentence[i]
		if v == '\n' || v == '\r' {
			if len(table[0]) > 0 {
				result += Print(table[:])
				table = [8][]string{}
			} else {
				result += "\n"
			}
			continue
		}

		for flag {
			if v == num {
				for j := 0; j < 8; j++ {
					table[j] = append(table[j], lines[counter])
					counter++
				}

				flag = false
			} else {
				num++
				counter += 9
			}
		}

		flag = true
		num = 32
		counter = 1
	}

	if len(table[0]) > 0 {
		result += Print(table[:])
	}

	return result
}
