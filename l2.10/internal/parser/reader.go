package parser

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile читает файл по заданному пути и возвращает слайс строк с содержимым.
func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return data, nil
}

// ReadStdin читает строки из стандартного ввода
func ReadStdin() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading stdin: %w", err)
	}
	return data, nil
}
