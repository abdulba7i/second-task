package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func stringUnpacking(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result string
	var escaped bool

	for i := 0; i < len(s); i++ {
		char := rune(s[i])

		if escaped {
			result += string(char)
			escaped = false
			continue
		}

		if char == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(char) {
			count, _ := strconv.Atoi(string(char))

			if len(result) == 0 {
				return "", errors.New("цифра не может быть первой в строке")
			}
			lastChar := rune(result[len(result)-1])

			if count > 1 {
				result += strings.Repeat(string(lastChar), count-1)
			}

			continue
		}

		if unicode.IsLetter(char) {
			result += string(char)
		} else {
			return "", errors.New("недопустимый символ")
		}
	}

	if escaped {
		return "", errors.New("строка заканчивается символом escape без следующего символа")
	}

	allDigits := true
	for _, r := range s {
		if !unicode.IsDigit(r) {
			allDigits = false
			break
		}
	}
	if allDigits {
		return "", errors.New("строка не может состоять только из цифр")
	}

	return result, nil
}

func main() {
	testsWord := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
	}

	for _, test := range testsWord {
		result, err := stringUnpacking(test)
		if err != nil {
			fmt.Printf("Строка: %q | Результат: %v\n", test, err)
		} else {
			fmt.Printf("Строка: %q | Результат: %q\n", test, result)
		}
	}
}
