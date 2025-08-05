package config

import (
	"fmt"
	"wb-tech-2/l2.10/internal/parser"

	"github.com/spf13/pflag"
)

// NewConfigFromFlags создает и возвращает конфигурацию,
// считывая параметры из флагов командной строки.
func NewConfigFromFlags() (*parser.Config, error) {
	// объявляем всевозможные пути сортировки:
	// -n, -r, -u, -M, -b, -c, -h
	column := pflag.IntP("key", "k", 1, "column to sort by (1-based index)")
	numeric := pflag.BoolP("numeric", "n", false, "sort numerically")
	reverse := pflag.BoolP("reverse", "r", false, "sort in reverse order")
	unique := pflag.BoolP("unique", "u", false, "output only unique lines")
	month := pflag.BoolP("month", "M", false, "sort by month name")
	ignoreTrailing := pflag.BoolP("ignore-trailing-blanks", "b", false, "ignore trailing blanks")
	checkSorted := pflag.BoolP("check", "c", false, "check if data is sorted")
	humanNumeric := pflag.BoolP("human-numeric", "h", false, "sort by human-readable numeric values (e.g., 2K, 1M)")

	pflag.Parse()

	// если номер указываемой колонки меньше 1, отдаем ошибку
	if *column < 1 {
		return nil, fmt.Errorf("column number must be >= 1: %d", *column)
	}

	// читаем путь к файла из аргумента и передаем в конфиг
	args := pflag.Args()
	inputFile := "-"
	if len(args) > 0 {
		inputFile = args[0]
	}

	// передаем в структуру параметры, переданные из пути
	config := parser.Config{
		Column:         *column,
		Numeric:        *numeric,
		Reverse:        *reverse,
		Unique:         *unique,
		Month:          *month,
		IgnoreTrailing: *ignoreTrailing,
		CheckSorted:    *checkSorted,
		HumanNumeric:   *humanNumeric,
		InputFile:      inputFile,
	}

	return &config, nil
}
