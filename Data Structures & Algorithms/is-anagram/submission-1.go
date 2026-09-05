func isAnagram(s string, t string) bool {
	sarr := make([]int, 26)

	for _, c := range s {
		ai := c - 'a'
		sarr[ai]++
	}

	for _, c := range t {
		ai := c - 'a'
		sarr[ai]--
	}

	for _, n := range sarr {
		if n != 0 {
			return false
		}
	}

	return true
}
