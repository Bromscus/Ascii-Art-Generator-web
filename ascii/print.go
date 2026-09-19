package ascii

func Print(result [][]string) string {
	str := ""
	for _, row := range result {
		for _, cell := range row {
			str += cell
		}
		str += "\n"
	}
	return str
}
