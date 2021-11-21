package X

func intToRoman(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM", "M"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	return thousands[(num/1000)%10] + hundreds[(num/100)%10] + tens[(num/10)%10] + ones[(num)%10]
}

