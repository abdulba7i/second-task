package cut

import "strings"

func CutLineReader(line string, delimiter string, fields map[int]struct{}, separatedOnly bool) (string, bool) {
	if separatedOnly && !strings.Contains(line, delimiter) {
		return "", false
	}

	columns := strings.Split(line, delimiter)
	var selected []string

	for i, col := range columns {
		if _, ok := fields[i+1]; ok {
			selected = append(selected, col)
		}
	}

	if len(selected) == 0 {
		return "", false
	}

	return strings.Join(selected, delimiter), true
}
