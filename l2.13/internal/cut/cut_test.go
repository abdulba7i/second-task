package cut

import (
	"testing"
)

func TestCutLines(t *testing.T) {
	lines := []string{
		"alpha,beta,gamma,delta",
		"one,two,three",
		"red,green,blue,yellow",
		"no_delimiter_here",
	}

	tests := []struct {
		name      string
		lines     []string
		delimiter rune
		fields    []int
		separated bool
		want      []string
	}{
		{
			name:      "basic fields",
			lines:     lines,
			delimiter: ',',
			fields:    []int{1, 3},
			separated: false,
			want:      []string{"alpha,gamma", "one,three", "red,blue", "no_delimiter_here"},
		},
		{
			name:      "separated flag filters no delimiter",
			lines:     lines,
			delimiter: ',',
			fields:    []int{2},
			separated: true,
			want:      []string{"beta", "two", "green"},
		},
		{
			name:      "out of range fields ignored",
			lines:     lines,
			delimiter: ',',
			fields:    []int{10},
			separated: false,
			want:      []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldsMap := toFieldMap(tt.fields)
			var got []string
			for _, line := range tt.lines {
				out, ok := CutLineReader(line, string(tt.delimiter), fieldsMap, tt.separated)
				if ok {
					got = append(got, out)
				}
			}

			if len(got) != len(tt.want) {
				t.Fatalf("got %d lines, want %d lines", len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("line %d: got %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func toFieldMap(fields []int) map[int]struct{} {
	m := make(map[int]struct{}, len(fields))
	for _, f := range fields {
		m[f] = struct{}{}
	}
	return m
}
