package _8XX

import "testing"

func Test_longestBeautifulSubstring(t *testing.T) {

	tests := []struct {
		name string
		word string
		want int
	}{{word :"aeiaaioaaaaeiiiiouuuooaauuaeiu", want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestBeautifulSubstring(tt.word); got != tt.want {
				t.Errorf("longestBeautifulSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
