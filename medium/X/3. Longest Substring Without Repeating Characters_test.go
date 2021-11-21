package X

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	v := lengthOfLongestSubstring(s)
	fmt.Println(v)

}
