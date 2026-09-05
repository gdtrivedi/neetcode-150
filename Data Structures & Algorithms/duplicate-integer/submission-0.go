func hasDuplicate(nums []int) bool {
	numsmap := make(map[int]bool, 0)

	for _, n := range nums {
		if _, ok := numsmap[n]; ok {
			// value exist
			return true
		}
		numsmap[n] = true
	}

	return false
}
