package X

func groupAnagrams(strs []string) [][]string {
	strMap := map[string][]string{}

	for _, s := range strs {
		keyRune := make([]rune, 26)
		for _, c := range s {
			keyRune[c-'a']++
		}
		key := string(keyRune)
		strMap[key] = append(strMap[key], s)
	}

	res := [][]string{}
	for _, v := range strMap {
		res = append(res, v)
	}

	return res
}
