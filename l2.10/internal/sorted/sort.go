package sorted

import (
	"fmt"
	"sort"
	"wb-tech-2/l2.10/internal/parser"
	"wb-tech-2/l2.10/internal/utils"
)

type sortableLine struct {
	original string
	key      string
}

// Sorted хранит конфигурацию структуры и переданные данные из файла
type Sorted struct {
	Config *parser.Config
	Data   []string
}

// NewSorted создает новый экземпляр Sorted
func NewSorted(config *parser.Config, data []string) *Sorted {
	return &Sorted{Config: config, Data: data}
}

// Sort выполняет сортировку данных в соответствии с конфигурацией и возвращает отсортированный слайс строк.
func (s *Sorted) Sort() ([]string, error) {
	var sortableLines []sortableLine

	// выделяем ключ для каждой строки
	for _, line := range s.Data {
		fields := utils.SplitLine(line)                  // здесь разбиваем строку на колонки
		key := utils.ExtractKey(fields, s.Config.Column) // вытаскиваем нужные столбец для параметра -k

		if s.Config.IgnoreTrailing {
			key = utils.TrimTrailingBlanks(key) // удаляем хвостовые пробелы (параметр -b)
		}

		sortableLines = append(sortableLines, sortableLine{
			original: line,
			key:      key,
		})
	}

	// сортируем срез по ключу, вместе с тем, передаем параметры сортировки s.Config
	sort.Slice(sortableLines, func(i, j int) bool {
		return compareLines(sortableLines[i], sortableLines[j], s.Config)
	})

	// Извлекаем отсортированные родные строки
	var result []string
	for _, line := range sortableLines {
		result = append(result, line.original)
	}

	// метод с исключением дубликатов в данных (параметр -u)
	if s.Config.Unique {
		result = utils.UniqueLines(result)
	}

	// проверка корректности сортировки (параметр -c)
	if s.Config.CheckSorted {
		for i := 1; i < len(sortableLines); i++ {
			if compareLines(sortableLines[i-1], sortableLines[i], s.Config) == false {
				return nil, fmt.Errorf("data is not sorted")
			}
		}
	}

	return result, nil
}

// Сравниваем два элемента с учётом выбранного режима сортировки
func compareLines(a, b sortableLine, cfg *parser.Config) bool {
	var result bool

	if cfg.Month { // преобразование месяца в число
		af := utils.MonthToNumber(a.key)
		bf := utils.MonthToNumber(b.key)
		result = af < bf
	} else if cfg.HumanNumeric { // преобразование числа с суффиксами (К, М, ...)
		af := utils.ParseHumanSize(a.key)
		bf := utils.ParseHumanSize(b.key)
		result = af < bf
	} else if cfg.Numeric { // преобразование обычных чисел
		af := utils.ParseNumeric(a.key)
		bf := utils.ParseNumeric(b.key)
		result = af < bf
	} else { // иначе обычное лексикографическое сравнение
		result = a.key < b.key
	}

	// переворачиваем строку, если параметр задан (параметр -r)
	if cfg.Reverse {
		return !result
	}

	return result
}
