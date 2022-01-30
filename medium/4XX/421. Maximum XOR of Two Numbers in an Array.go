package _XX

type trie struct {
	isLeaf bool
	bit    int
	value  int
	bitPos int
	child  [2]*trie
}

func findMaximumXOR(nums []int) int {
	root := new(trie)
	mask := 0x80000000
	//fmt.Printf("%b\n", mask)
	//fmt.Println(strconv.FormatInt(int64(mask), 2)) // 1111011
	//fmt.Printf("%b", mask)
	for _, n := range nums {
		node := root
		temp := n

		//fmt.Printf("%b", n)
		for i := 0; i < 32; i++ {
			//fmt.Println(strconv.FormatInt(int64(temp), 2)) // 1111011
			if i == 31 {
				node.value = n
				node.isLeaf = true
				if temp&mask == 0 {
					node.bit = 0
				} else {
					node.bit = 1
				}

				//fmt.Printf("%d", node.bit)
				break
			}

			if temp&mask == 0 {
				//case 0
				//fmt.Printf("0")
				if node.child[0] == nil {
					node.child[0] = new(trie)
					node.child[0].bit = 0
					node.bitPos = i + 1
				}
				node = node.child[0]
			} else {
				//case 1
				//fmt.Printf("1")
				if node.child[1] == nil {
					node.child[1] = new(trie)
					node.child[1].bit = 1
					node.bitPos = i + 1

				}
				node = node.child[1]
			}
			temp = temp << 1
		}
		//fmt.Println()
	}

	firstMap, secondMap := map[int]*trie{}, map[int]*trie{}
	nFirstMap, nSecondMap := map[int]*trie{}, map[int]*trie{}
	res := []*trie{}
	for _, n := range root.child {
		if n == nil {
			continue
		}

		if n.bit == 0 {
			firstMap[n.bit] = n
		} else {
			secondMap[n.bit] = n
		}
	}

	for len(firstMap) > 0 || len(secondMap) > 0 {
		// 2 map set case

		if len(res) == 2 && res[0].isLeaf && res[1].isLeaf {
			return res[0].value ^ res[1].value
		}
		res = []*trie{}

		if len(firstMap) > 0 && len(secondMap) > 0 {
			getSameMap(firstMap, nFirstMap)
			getSameMap(secondMap, nSecondMap)

			res = []*trie{}

			res = removeDuplicates(nFirstMap, nSecondMap, res)

			firstMap = nFirstMap
			secondMap = nSecondMap
			nFirstMap = make(map[int]*trie)
			nSecondMap = make(map[int]*trie)

		} else
		// one map set case
		{
			res = getDiffMap(firstMap, res, nFirstMap, nSecondMap)
			res2 := getDiffMap(secondMap, res, nFirstMap, nSecondMap)
			res = append(res, res2...)

			firstMap = nFirstMap
			secondMap = nSecondMap
			nFirstMap = make(map[int]*trie)
			nSecondMap = make(map[int]*trie)

			continue
		}

	}

	return 0
}

func removeDuplicates(nFirstMap map[int]*trie, nSecondMap map[int]*trie, res []*trie) []*trie {
	for _, n := range nFirstMap {
		if _, ok := nSecondMap[getComplement(n.bit)]; !(len(nFirstMap) == 1) && !ok {
			delete(nFirstMap, n.bit)
		} else {
			res = append(res, n)

		}
	}

	for _, n := range nSecondMap {
		if _, ok := nFirstMap[getComplement(n.bit)]; !(len(nSecondMap) == 1) && !ok {
			delete(nSecondMap, n.bit)
		} else {
			res = append(res, n)
		}
	}
	return res
}

func getDiffMap(firstMap map[int]*trie, res []*trie, nFirstMap map[int]*trie, nSecondMap map[int]*trie) []*trie {
	for _, n := range firstMap {
		if n == nil {
			continue
		}

		for _, c := range n.child {
			if c == nil {
				continue
			}

			res = append(res, c)
			if c.bit == 0 {
				nFirstMap[c.bit] = c
			} else {
				nSecondMap[c.bit] = c
			}
		}
	}
	return res
}

func getSameMap(firstMap map[int]*trie, nFirstMap map[int]*trie) {
	for _, n := range firstMap {
		if n == nil {
			continue
		}

		for _, c := range n.child {
			if c == nil {
				continue
			}
			nFirstMap[c.bit] = c
		}
	}
}

func getComplement(val int) int {
	if val == 0 {
		return 1
	}

	return 0
}
