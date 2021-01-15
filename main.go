package main

import (
	"fmt"
	"yk.com/yugencoder/LC/medium/13XX"
)

func main() {
	n := 4
	left := []int{1, -1, 3, -1}
	right := []int{2, -1, -1, -1}
	res := t13XX.ValidateBinaryTreeNodes(n,left, right)
	fmt.Print(res)
}
