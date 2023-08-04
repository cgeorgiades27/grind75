package validanagram

func ValidAnagramCG(s, t string) bool {
	var maxSizeS, maxSizeT int32
	for _, r := range s {
		if r > maxSizeS {
			maxSizeS = r
		}
	}

	for _, r := range t {
		if r > maxSizeT {
			maxSizeT = r
		}
	}

	if maxSizeS != maxSizeT {
		return false
	}

	m1 := make([]rune, maxSizeS+1)
	for _, elem := range s {
		m1[elem]++
	}

	m2 := make([]rune, maxSizeS+1)
	for _, elem := range t {
		m2[elem]++
	}

	for i, r := range m1 {
		if m2[i] != r {
			return false
		}
	}

	return true
}
