package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	anagrams := searchAnagrams(input)

	for key, value := range anagrams {
		fmt.Println(key, value)
	}
}

func searchAnagrams(arr []string) map[string][]string {
	groupsAnagram := make(map[string][]string)
	result := make(map[string][]string)

	for _, word := range arr {
		lWord := strings.ToLower(word)

		key := sortedLetters(lWord)

		groupsAnagram[key] = append(groupsAnagram[key], lWord)
	}

	for _, anagram := range groupsAnagram {
		if len(anagram) > 1 {
			sort.Strings(anagram)
			result[anagram[0]] = anagram
		}
	}

	return result
}

func sortedLetters(word string) string {
	letters := []rune(word)
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	return string(letters)
}
