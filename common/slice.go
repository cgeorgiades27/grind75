package common

func CheckSliceEquality[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, e := range s1 {
		if s2[i] != e {
			return false
		}
	}
	return true
}
