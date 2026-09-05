func twoSum(nums []int, target int) []int {
    m := make(map[int]int, 0)

	for i, n := range nums {
		sub := target - n
		if val, ok := m[sub]; ok {
			return []int{val, i}
		}
		m[n] = i
	}

	return []int{-1, -1}
}
