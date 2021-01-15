package _14XX

import "fmt"

func getFolderNames(names []string) []string {
	var res []string
	nameMap := map[string]int{}

	for _, name := range names {
		k := nameMap[name]
		if k == 0 {
			nameMap[name] = 1
			res = append(res, name)

			continue
		}

		newName := fmt.Sprintf("%s(%d)", name, k)
		k++

		for ; nameMap[newName] > 0; {
			newName = fmt.Sprintf("%s(%d)", name, k)
			k++
		}

		nameMap[name] = k
		nameMap[newName] = 1
		res = append(res, newName)
	}

	return res
}
