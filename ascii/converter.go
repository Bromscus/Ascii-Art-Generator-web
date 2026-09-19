package ascii

import "strings"

func converter(arg []string) []rune {
	text := strings.Join(arg, " ")
	return []rune(text)
}
