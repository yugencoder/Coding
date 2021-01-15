package t13XX

import (
	"fmt"
	"testing"
)

func Test_breakPalindrome(t *testing.T) {
	tests := []struct {
		name       string
		palindrome string
	}{
		{name: "aa", palindrome: "abccba"},
		{
			name:       "",
			palindrome: "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := breakPalindrome(tt.palindrome)
			fmt.Println(got)
		})
	}


}
