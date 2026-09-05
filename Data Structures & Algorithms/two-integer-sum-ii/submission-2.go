func twoSum(numbers []int, target int) []int {
	lp := 0
	rp := len(numbers) - 1

	for lp < rp {
		sum := numbers[lp] + numbers[rp]
		if sum == target {
			return []int{lp + 1, rp + 1}
		}

		if sum < target {
			lp++
		}

		if sum > target {
			rp--
		}
	}
	return []int{}
}
