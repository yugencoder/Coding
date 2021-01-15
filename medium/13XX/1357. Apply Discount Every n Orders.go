package t13XX

type Cashier struct {
	currentCustomerCount int
	n, discount          int
	priceByProductID     map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	priceByProductID := map[int]int{}

	for i := range products {
		priceByProductID[products[i]] = prices[i]
	}

	return Cashier{
		currentCustomerCount: 0,
		n:                    n,
		discount:             discount,
		priceByProductID:     priceByProductID,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	x := float64(0)
	this.currentCustomerCount++

	for i := range product {
		x += float64(this.priceByProductID[product[i]] * amount[i])
	}

	if this.currentCustomerCount == this.n {
		this.currentCustomerCount = 0

		return x - float64(this.discount)*x/100
	}

	return x
}
