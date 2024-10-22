package str

import "strings"

// a: "101" b: "1011"
//
//	101
//
// 1011
// Output: "10000"
func AddBinary(a, b string) string {
	res := strings.Builder{}

	la, lb := len(a), len(b)
	carry := 0
	for i, j := la-1, lb-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(a[i] - '0')
		}
		if j >= 0 {
			y = int(b[j] - '0')
		}

		sum := x + y + carry

		res.WriteByte(byte(sum%2 + '0'))
		carry = sum / 2
	}

	if carry == 1 {
		res.WriteByte('1')
	}

	reverse := func(s string) string {
		bs := []byte(s)
		for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
			bs[i], bs[j] = bs[j], bs[i]
		}

		return string(bs)
	}

	return reverse(res.String())
}
