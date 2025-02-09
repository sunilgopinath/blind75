package blind75

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(s); i++ {
		odd := expandFromCenter(s, i, i)
		if len(odd) > len(res) {
			res = odd
		}
		even := expandFromCenter(s, i, i+1)
		if len(even) > len(res) {
			res = even
		}
	}
	return res
}

func expandFromCenter(s string, l, r int) string {
	for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
