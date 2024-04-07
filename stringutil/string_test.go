package stringutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
	}

	for _, test := range tests {
		if output := Reverse(test.input); output != test.expected {
			t.Errorf("Reverse(%s) = %s; want %s", test.input, output, test.expected)
		}
	}
}

func TestSymbolCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"world", 5},
		{"", 0},
		{"kj", 2},
	}

	for _, test := range tests {
		if output := SymbolCount(test.input); output != test.expected {
			t.Errorf("SymbolCount(%s) = %d; want %d", test.input, output, test.expected)
		}
	}
}
