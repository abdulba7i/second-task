package config

// Config содержит конфигурацию программы.
type Config struct {
	After      int
	Before     int
	Context    int
	CountOnly  bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Pattern    string
}
