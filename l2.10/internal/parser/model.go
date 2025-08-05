package parser

// Config структура с полями параметров сортировки
type Config struct {
	Column         int
	Numeric        bool
	Reverse        bool
	Unique         bool
	Month          bool
	IgnoreTrailing bool
	CheckSorted    bool
	HumanNumeric   bool
	InputFile      string
}
