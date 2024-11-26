package str

// Input: str="abcdefg" ,n=2
// Output: fgabcde

// 1) 整体翻转: gfedcba
// 2) 局部翻转: fgabcde

func reverse(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func RotateString(s string, n int) string {
	n = n % len(s)
	if n > 0 {
		bs := []byte(s)
		reverse(bs, 0, len(s)-1)
		reverse(bs, 0, n-1)
		reverse(bs, n, len(s)-1)

		return string(bs)
	}

	return s
}

// "abcdedfg"
// 1) 局部翻转: edcbagf
// 2) 整体翻转: fgabcde
func RotateStringII(s string, n int) string {
	if n == len(s) {
		return s
	}

	if n > len(s) {
		n = n % len(s)
	}

	bs := []byte(s)
	reverse(bs, 0, len(bs)-n-1)
	reverse(bs, len(bs)-n, len(bs)-1)
	reverse(bs, 0, len(bs)-1)

	return string(bs)
}
