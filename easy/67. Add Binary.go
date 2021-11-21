package easy

import (
	"fmt"
	"strconv"
	"strings"
)

func AddBinary(a string, b string) string {
	n1 := strings.Split(a,"")
	n2 := strings.Split(b,"")

	n1 = reverse(n1)
	n2 = reverse(n2)

	sum := make([]int, max(len(a),len(b)))
	c := 0


	for i:=0; i< len(sum);i++{
		aVal,bVal := 0,0

		if i<len(n1){
			aVal,_ = strconv.Atoi(n1[i])
		}
		if i<len(n2){
			bVal,_ = strconv.Atoi(n2[i])
		}


		res := aVal + bVal + c
		if res > 1{
			c = 1
			sum[i] = 0
			if res > 2{
				sum[i] = 1
			}
		}else{
			sum[i] = res
			c = 0
		}
	}

	if c == 1{
		sum = append(sum,c)
	}

	var ansStr string
	for _, a:= range sum{
		ansStr = fmt.Sprintf("%v%v",a,ansStr)
	}

	return ansStr
}


func reverse(a []string) []string{
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return a
}



func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}


func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}