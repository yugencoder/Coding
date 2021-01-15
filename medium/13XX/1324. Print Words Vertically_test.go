package t13XX

import (
	"fmt"
	"testing"
)

func Test_printVertically(t *testing.T) {

	s := "CONTEST IS COMING"

	got := printVertically(s)
	fmt.Printf("%+v", got)
}
