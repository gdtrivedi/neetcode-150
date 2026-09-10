func lengthOfLongestSubstring(s string) int {
	ptr := 0
	tmpStr := ""
	l := len(s)
	longSubStrLen := 0
	for ptr < l {
		rc := rune(s[ptr])
		rci := strings.IndexRune(tmpStr, rc)

		// found character
		if rci != -1 {
			tmpStr = tmpStr[rci+1:]
		}

		tmpStr = tmpStr + string(rc)
		longSubStrLen = max(longSubStrLen, len(tmpStr))
		ptr += 1
	}

	return longSubStrLen
}
