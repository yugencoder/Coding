package t13XX

type CustomStack struct {
	arr []int
	top int
}

//func Constructor(maxSize int) CustomStack {
//	return CustomStack{
//		make([]int, maxSize),
//		0,
//	}
//}

func (this *CustomStack) Push(x int) {
	this.top++
	this.arr[this.top] = x
}

func (this *CustomStack) Pop() int {
	temp := this.arr[this.top]
	this.top--
	this.arr = this.arr[:this.top]
	return temp

}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k; i++ {
		this.arr[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
