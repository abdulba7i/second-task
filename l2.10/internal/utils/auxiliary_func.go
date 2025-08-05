package utils

import (
	"strconv"
	"strings"
)

// SplitLine столбцы по табуляции
func SplitLine(line string) []string {
	return strings.Split(line, "\t")
}

// ExtractKey извлекает ключ по номеру столбца
func ExtractKey(fields []string, col int) string {
	if col-1 >= 0 && col-1 < len(fields) {
		return fields[col-1]
	}
	return ""
}

// ParseNumeric преобразование строки в float, для нумерации
func ParseNumeric(val string) float64 {
	val = strings.TrimSpace(val)
	if val == "" {
		return 0
	}
	f64, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}

	return f64
}

// UniqueLines возвращает слайс с уникальными строками из исходного слайса
func UniqueLines(lines []string) []string {
	duplicateMap := map[string]bool{}
	var result []string

	for _, line := range lines {
		if !duplicateMap[line] {
			duplicateMap[line] = true
			result = append(result, line)
		}
		// если уже есть — пропускаем
	}

	return result
}

// MonthToNumber преобразование название месяца в нумерацию
func MonthToNumber(month string) int {
	months := map[string]int{
		"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6,
		"Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
	}
	return months[strings.Title(strings.TrimSpace(month))]
}

// ParseHumanSize преобразование строки с суффиксами (1K, 2M, 3G, 4T) в float64
func ParseHumanSize(val string) float64 {
	val = strings.TrimSpace(val)
	if val == "" {
		return 0
	}
	mult := 1.0
	unit := val[len(val)-1]
	if unit < '0' || unit > '9' {
		switch unit {
		case 'K', 'k':
			mult = 1e3
		case 'M', 'm':
			mult = 1e6
		case 'G', 'g':
			mult = 1e9
		case 'T', 't':
			mult = 1e12
		}
		val = val[:len(val)-1]
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}
	return f * mult
}

// TrimTrailingBlanks удаляет хвостовые пробелы
func TrimTrailingBlanks(s string) string {
	return strings.TrimRight(s, " \t")
}
