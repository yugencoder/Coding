package t13XX

import (
	"fmt"
	"testing"
)

func Test_minFlips(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
	}{
		{
			name: "",
			a:    2,
			b:    6,
			c:    5,
		},
		//{
		//	name: "",
		//	a:    4,
		//	b:    2,
		//	c:    7,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minFlips(tt.a, tt.b, tt.c)
			fmt.Println(got)
		})
	}
}
