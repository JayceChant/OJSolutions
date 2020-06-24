package main

var matchMap = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValidMap(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s))
	for _, ch := range s {
		match, ok := matchMap[ch]
		if ok {
			if len(stack) == 0 || match != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var match rune
	stack := make([]rune, 0, len(s)/2)
	for _, char := range s {
		isRight := true
		switch char {
		case ')':
			match = '('
		case ']':
			match = '['
		case '}':
			match = '{'
		default:
			isRight = false
		}
		if isRight {
			if len(stack) == 0 || match != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if len(stack)*2 == len(s) {
			return false
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
