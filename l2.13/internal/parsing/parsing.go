package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseFields(expr string) (map[int]struct{}, error) {
	fields := make(map[int]struct{})
	parts := strings.Split(expr, ",")

	for _, part := range parts {
		if part == "" {
			return nil, fmt.Errorf("empty field")
		}
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			start, err1 := strconv.Atoi(rangeParts[0])
			end, err2 := strconv.Atoi(rangeParts[1])
			if err1 != nil || err2 != nil || start > end || start < 1 {
				return nil, fmt.Errorf("invalid field range: %s", part)
			}

			for i := start; i <= end; i++ {
				fields[i] = struct{}{}
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil || num < 1 {
				return nil, fmt.Errorf("invalid field: %s", part)
			}
			fields[num] = struct{}{}
		}
	}

	return fields, nil
}

// FieldsMapToSlice функция для тестов
func FieldsMapToSlice(fields map[int]struct{}) []int {
	res := make([]int, 0, len(fields))
	for k := range fields {
		res = append(res, k)
	}

	if len(res) > 1 {
		for i := 0; i < len(res)-1; i++ {
			for j := i + 1; j < len(res); j++ {
				if res[i] > res[j] {
					res[i], res[j] = res[j], res[i]
				}
			}
		}
	}
	return res
}
