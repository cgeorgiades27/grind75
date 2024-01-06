package twosum

func twoSumCg(target int, nums []int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
		m[num-target] = i
	}
	return []int{}
}
