package blind75

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	maxSeq, left := 0, 0
	for right := 0; right < len(s); right++ {
		for m[s[right]] {
			delete(m, s[left])
			left++
		}
		m[s[right]] = true
		maxSeq = max(maxSeq, right-left+1)

	}
	return maxSeq
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
