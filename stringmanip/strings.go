package stringmanip

// Substring provides a safe way to extract a substring from a UTF-8 string.
// From this discussion:
//   https://groups.google.com/d/msg/golang-nuts/cGq1Irv_5Vs/0SKoj49BsWQJ
func Substring(s string, p, l int) string {
	if p < 0 || l <= 0 {
		return ""
	}
	c := []rune(s)
	if p > len(c) {
		return ""
	} else if p+l > len(c) || p+l < p {
		return string(c[p:])
	}
	return string(c[p : p+l])
}
