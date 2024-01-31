package validparentheses

func validParen(s string) bool {
	m := make([]int, 128)
	m['}'] = '{'
	m[')'] = '('
	m[']'] = '['

	stack := make([]byte, len(s))
	sPtr := 0
	for i := 0; i < len(s); i++ {
		for i < len(s) &&
			s[i] != '}' &&
			s[i] != ']' &&
			s[i] != ')' {
			stack[sPtr] = s[i]
			sPtr++
			i++
		}

		if i == len(s) {
			return false
		}

		sPtr--
		if m[s[i]] != int(stack[sPtr]) {
			return false
		}
	}

	return sPtr == 0
}
