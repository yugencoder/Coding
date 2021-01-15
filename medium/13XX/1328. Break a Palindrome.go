package t13XX

func breakPalindrome(palindrome string) string {
	n := len(palindrome) / 2
	replaced := false
	for i := 0; i < n; i++ {
		if palindrome[i] != 'a' {

			palindrome = palindrome[0:i] + "a" + palindrome[i+1:]
			replaced = true
			break
		}
	}

	if !replaced {
		for i := 0; i < len(palindrome)/2; i++ {
			if palindrome[i] != 'b' {
				palindrome = palindrome[0:i] + "b" + palindrome[i+1:]
				replaced = true
				palindrome = Reverse(palindrome)
				break
			}
		}
	}

	if !replaced {
		return ""
	}

	return palindrome
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
