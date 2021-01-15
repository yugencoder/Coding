package t13XX

type ProductOfNumbers struct {
	prod     []int
	lastZero int
}

func ProductOfNumbersConstructor() ProductOfNumbers {
	productOfNumbers := ProductOfNumbers{}
	productOfNumbers.prod = make([]int, 0)

	return productOfNumbers
}

func (this *ProductOfNumbers) Add(num int) {
	n := len(this.prod)

	if num == 0 {
		this.prod = make([]int, 0)
		return
	}

	if n == 0 {
		this.prod = append(this.prod, num)
		return
	}

	this.prod = append(this.prod, this.prod[n-1]*num)

}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.prod)

	if k > n {
		return 0
	}

	if n == k {
		return this.prod[n-1]
	}

	return this.prod[n-1] / this.prod[n-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
