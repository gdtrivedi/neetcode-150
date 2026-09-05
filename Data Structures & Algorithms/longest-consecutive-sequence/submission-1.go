func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	res := 1

	store := make(map[int]struct{})
	for _, n := range nums {
		store[n] = struct{}{}
	}

	for _, n := range nums {
		streak := 0
		cur := n
		_, ok := store[cur]
		for ok {
			streak++
			cur++
			res = max(res, streak)
			_, ok = store[cur]
		}
	}

	return res
}
