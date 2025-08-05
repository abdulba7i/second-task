package main

import (
	"fmt"
	"log"
	"wb-tech-2/l2.10/internal/config"
	"wb-tech-2/l2.10/internal/parser"
	"wb-tech-2/l2.10/internal/sorted"
)

func main() {
	// собираем конфиг
	cfg, err := config.NewConfigFromFlags()
	if err != nil {
		log.Fatalf("wrong parse config: %v", err)
	}

	// чтение файла и добавление данных в слайс data
	var data []string
	if cfg.InputFile == "-" {
		data, err = parser.ReadStdin()
	} else {
		data, err = parser.ReadFile(cfg.InputFile)
	}
	if err != nil {
		log.Fatalf("error read file: %v", err)
	}

	// основная логика сортировки
	sorter := sorted.NewSorted(cfg, data)
	sortedLines, err := sorter.Sort()
	if err != nil {
		log.Fatalf("sort error: %v", err)
	}

	// вывод результата
	for _, line := range sortedLines {
		fmt.Println(line)
	}
}
