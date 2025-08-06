package parsing

import (
	"reflect"
	"testing"
)

func TestParseFields(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
		err      bool
	}{
		{"1,3-5", []int{1, 3, 4, 5}, false},
		{"2", []int{2}, false},
		{"4-4", []int{4}, false},
		{"3-1", nil, true}, // кейс с обратным диапазоном, должна вернуть ошибку ошибка
		{"a,2", nil, true}, // кейс с неверным номером
		{"", nil, true},    // кейс с пустой строкой, должна вернуть ошибка
		{"1,2,5-7", []int{1, 2, 5, 6, 7}, false},
	}

	for _, tt := range tests {
		got, err := ParseFields(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("ParseFields(%q) error = %v, wantErr %v", tt.input, err, tt.err)
			continue
		}
		if !tt.err {
			gotSlice := FieldsMapToSlice(got)
			if !reflect.DeepEqual(gotSlice, tt.expected) {
				t.Errorf("ParseFields(%q) = %v, want %v", tt.input, gotSlice, tt.expected)
			}
		}
	}
}
