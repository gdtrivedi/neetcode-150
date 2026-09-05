func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	m := make(map[string][]string, 0)
	for _, str := range strs {
		sortedstr := sortWordByCharacter(str)
		m[sortedstr] = append(m[sortedstr], str)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func sortWordByCharacter(w string) string {
	// Convert the string to a slice of runes
	runes := []rune(w)

	// Use sort.Slice with a custom less function
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j] // Sort in ascending order
	})

	// Convert the sorted slice of runes back to a string
	return string(runes)
}
