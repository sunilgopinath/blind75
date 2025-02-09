package blind75

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(s); i++ {
		res += expandFromCenterCounter(s, i, i)
		res += expandFromCenterCounter(s, i, i+1)
	}
	return res
}

func expandFromCenterCounter(s string, l, r int) int {
	res := 0
	for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}
