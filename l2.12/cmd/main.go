package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"wb-tech-2/l2.12/internal/config"
	"wb-tech-2/l2.12/internal/grep"
	"wb-tech-2/l2.12/internal/reader"
)

func main() {
	// прописываем все необходимые флаги
	after := flag.Int("A", 0, "печатать N строк после совпадения")
	before := flag.Int("B", 0, "печатать N строк до совпадения")
	context := flag.Int("C", 0, "печатать N строк вокруг совпадения (до и после)")
	count := flag.Bool("c", false, "печатает только количество совпадающих строк")
	ignoreCase := flag.Bool("i", false, "игнорирует регистр")
	invert := flag.Bool("v", false, "выводит строки, не совпадающие с шаблоном")
	fixed := flag.Bool("F", false, "воспринимает шаблон как фиксированную строку")
	lineNum := flag.Bool("n", false, "выводит номера строк")

	flag.Parse()

	// Проверка: должен быть шаблон и путь к файлу
	args := flag.Args()
	if len(args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	pattern := args[0]
	filename := args[1]

	// чтение файла
	lines, err := reader.ReadLines(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	cfg := config.Config{
		After:      *after,
		Before:     *before,
		Context:    *context,
		CountOnly:  *count,
		IgnoreCase: *ignoreCase,
		Invert:     *invert,
		Fixed:      *fixed,
		LineNum:    *lineNum,
		Pattern:    pattern,
	}

	// основная логика работы утилиты grep
	result := grep.ProcessLines(lines, cfg)

	// вывод результата
	for _, line := range result {
		fmt.Println(line)
	}
}
