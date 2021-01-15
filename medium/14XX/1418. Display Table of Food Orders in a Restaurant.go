package _14XX

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	var items []string
	var res [][]string

	itemMap := map[string]bool{}
	tableOrderMap := map[string]map[string]int{}

	for _, o := range orders {
		tableNo := o[1]
		item := o[2]
		itemMap[item] = true

		if _, ok := tableOrderMap[tableNo]; !ok {
			tableOrderMap[tableNo] = map[string]int{}
		}

		tableOrderMap[tableNo][item] += 1
	}

	for k := range itemMap {
		items = append(items, k)
	}

	sort.Strings(items)

	for i := 0; i <= 500; i++ {
		var row []string
		iStr := strconv.Itoa(i)

		if i == 0 {
			row = append(row, "table")
			for _, it := range items {
				row = append(row, it)
			}
			continue
		}

		if tableOrder, ok := tableOrderMap[iStr]; ok {
			row = append(row, iStr)
			for _, it := range items {
				if _, ok := tableOrder[it]; ok {
					row = append(row, strconv.Itoa(tableOrder[it]))
				} else {
					row = append(row, "0")
				}
			}
			res = append(res, row)
		}
	}

	return res
}
