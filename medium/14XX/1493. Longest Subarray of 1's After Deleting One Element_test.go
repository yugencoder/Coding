package _14XX

import (
	"fmt"
	"testing"
)

func Test_longestSubarray2(t *testing.T) {

	nums := []int{0,1,1,1,0,1,1,0,1}
	nums = []int{1,1,0,0,1,1,1,0,1}
	fmt.Println(longestSubarray2(nums))


}
