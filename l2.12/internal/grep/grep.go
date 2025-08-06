package grep

import (
	"fmt"
	"regexp"
	"strings"
	"wb-tech-2/l2.12/internal/config"
)

// ProcessLines работает со строками согласно конфигурации
func ProcessLines(lines []string, cfg config.Config) []string {
	if cfg.Context > 0 {
		cfg.After = cfg.Context
		cfg.Before = cfg.Context
	}

	matches := make([]bool, len(lines))
	var re *regexp.Regexp
	var err error

	// Подготовка основного паттерна для поиска
	pattern := cfg.Pattern
	if cfg.IgnoreCase {
		pattern = "(?i)" + pattern
	}

	if !cfg.Fixed {
		re, err = regexp.Compile(pattern)
		if err != nil {
			return []string{fmt.Sprintf("regex compilation error: %v", err)}
		}
	}

	// Поиск совпадений
	for i, line := range lines {
		match := false
		if cfg.Fixed {
			content := line
			if cfg.IgnoreCase {
				content = strings.ToLower(content)
				pattern = strings.ToLower(cfg.Pattern)
			}
			match = strings.Contains(content, pattern)
		} else {
			match = re.MatchString(line)
		}
		if cfg.Invert {
			match = !match
		}
		matches[i] = match
	}

	// Если нужно сделать только подсчёт
	if cfg.CountOnly {
		count := 0
		for _, m := range matches {
			if m {
				count++
			}
		}
		return []string{fmt.Sprintf("%d", count)}
	}

	// Основной вывод с контекстом
	seen := make(map[int]bool)
	var result []string

	for i, matched := range matches {
		if !matched {
			continue
		}
		start := i - cfg.Before
		if start < 0 {
			start = 0
		}
		end := i + cfg.After
		if end >= len(lines) {
			end = len(lines) - 1
		}
		for j := start; j <= end; j++ {
			if seen[j] {
				continue
			}
			seen[j] = true
			line := lines[j]
			if cfg.LineNum {
				line = fmt.Sprintf("%d:%s", j+1, line)
			}
			result = append(result, line)
		}
	}

	return result
}
