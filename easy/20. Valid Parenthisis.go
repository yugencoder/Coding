package easy

func IsValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = push(stack, c)
		case ')', '}', ']':
			val := pop(&stack)
			if val != getComplement(c) {
				return false
			}
		}
	}

	return len(stack) == 0
}

func push(stack []rune, val rune) []rune {
	return append([]rune{val}, stack...)
}

func pop(stack *[]rune) rune {
	st := *stack
	if len(st) == 0 {
		return '!'
	}
	val := st[0]
	st = st[1:]

	*stack = st

	return val
}

func getComplement(val rune) rune {
	switch val {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	}

	return '!'
}
