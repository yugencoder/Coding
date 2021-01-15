package _14XX

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	arrMap := map[int]int{}
	type item struct {
		key int
		val int
	}

	itemArr := []item{}
	for _, a := range arr {
		arrMap[a] += 1
	}

	for k, v := range arrMap {
		itemArr = append(itemArr, item{
			key: k,
			val: v,
		})
	}

	sort.Slice(itemArr, func(i, j int) bool {
		return itemArr[i].val < itemArr[j].val
	})

	for _, item := range itemArr {
		if k <= 0 {
			break
		}

		arrMap[item.key] -= min(item.val, k)
		k -= min(item.val, k)

		if arrMap[item.key] == 0 {
			delete(arrMap, item.key)
		}
	}

	return len(arrMap)
}
