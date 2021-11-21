package _8XX

//func getNumberOfBacklogOrders(orders [][]int) int {
//	buyOrders := make([][]int, 0)
//	sellOrders := make([][]int, 0)
//
//	for _, o := range orders {
//		if o[2] == 0 {
//			buyOrders = insertInSortedArr(o, buyOrders)
//		} else {
//			sellOrders = insertInSortedArr(o, sellOrders)
//		}
//
//		for ; ; {
//			index, order := findOrdersLessThanEq(o, sellOrders)
//			if index == -1 {
//				break
//			}
//
//			fulfilled := min(order[1], o[1])
//			o[1] = o[1] - fulfilled
//
//			if fulfilled == order[1] {
//				sellOrders = append(sellOrders[:index], sellOrders[index+1:]...)
//			} else {
//				sellOrders[index][1] = order[1] - fulfilled
//			}
//			if o[1] == 0 {
//				break
//			}
//		}
//	}
//
//	res := 0
//	for _, o := range sellOrders {
//		res = (res + o[1]) % 1000000007
//	}
//
//	for _, o := range buyOrders {
//		res = (res + o[1]) % 1000000007
//	}
//
//	return res
//}
//
//
//func findOrdersLessThanEq(o []int, orders [][]int) (int, []int) {
//	for i := len(orders) - 1; i >= 0; i-- {
//		if orders[i][0] <= o[0] {
//			return i, orders[i]
//		}
//	}
//
//	return -1, nil
//}
//
//func insertInSortedArr(o []int, orders [][]int) [][]int {
//	for i, order := range orders {
//		if o[0] == order[0] {
//			order[1] += o[0]
//
//			return orders
//		}
//
//		if o[0] < order[0] {
//			if i == 0 {
//				return append([][]int{o}, orders...)
//			} else {
//				orders = append(orders, []int{})
//				copy(orders[i+1:], orders[i:])
//				orders[i] = o
//				return orders
//			}
//
//			return append(append(orders[:i], o), orders[i:]...)
//		}
//	}
//
//	return append(orders, o)
//}
