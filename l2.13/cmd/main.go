package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"wb-tech-2/l2.13/internal/cut"
	"wb-tech-2/l2.13/internal/parsing"
)

func main() {
	// Флаги
	fieldsFlag := flag.String("f", "", "список полей для вывода (например, 1,3-5)")
	delim := flag.String("d", "\t", "разделитель (по умолчанию табуляция)")
	separated := flag.Bool("s", false, "выводить только строки с разделителем")

	flag.Parse()

	if *fieldsFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: -f flag (fields) is required")
		os.Exit(1)
	}

	fields, err := parsing.ParseFields(*fieldsFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing fields: %v\n", err)
		os.Exit(1)
	}

	var input *os.File
	if len(flag.Args()) > 0 {
		file, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		result, ok := cut.CutLineReader(line, *delim, fields, *separated)
		if ok {
			fmt.Println(result)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
