package twosum

func twoSumCg(target int, nums []int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return []int{}
}
