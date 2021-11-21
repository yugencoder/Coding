package main

import (
	"fmt"
	"yk.com/yugencoder/LC/easy"
)

func main() {
	//l5 := new(easy.ListNode)
	//l1 := new(easy.ListNode)
	//
	//l5.Val = 5
	//l1.Val = 1
	//
	//l9 := new(easy.ListNode)
	//l9.Val = 9
	//
	//l8 := new(easy.ListNode)
	//l8.Val = 8
	//
	//l4 := new(easy.ListNode)
	//l4.Val = 4
	//
	//// 1->8->4
	//l1.Next = l8
	//l8.Next = l4
	//
	//// 5-> 9 -> 8 -> 4
	//l5.Next = l9
	//l9.Next = l8

//["TwoSum","add","add","add","find","find"]
	//[[],[1],[3],[5],[4],[7]]

	res := easy.TwoSum{}
	res.Add(1)
	res.Add(-1)
	//res.Add(5)
	//fmt.Println(res.Find(4))
	fmt.Println(res.Find(0))

	fmt.Println(res)
}
