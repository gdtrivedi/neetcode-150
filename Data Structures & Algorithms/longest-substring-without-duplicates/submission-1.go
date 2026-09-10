func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	maxLen, start := 0, 0

	for end, char := range s {
		// If the character was seen inside the current window, move the start
		if idx, duplicate := charMap[char]; duplicate && idx >= start {
			start = idx + 1
		}

		// Update/Insert the character's newest index
		charMap[char] = end

		// Calculate window size
		if currentLen := end - start + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
