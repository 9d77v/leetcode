package algorithm

//IsPalidrome 判断是否回文子串
func IsPalidrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
