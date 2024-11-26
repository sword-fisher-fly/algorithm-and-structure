package special

import "strings"

func convert(s string) int {
	num := 0
	for i := range s {
		num = num*10 + int(s[i]-'0')
	}

	return num
}

func CompareVersion(version1 string, version2 string) int {
	// write code here
	vv1 := strings.Split(version1, ".")
	vv2 := strings.Split(version2, ".")

	for i, j := 0, 0; i < len(vv1) && j < len(vv2); i, j = i+1, j+1 {
		v1, v2 := convert(vv1[i]), convert(vv2[j])
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}

	if len(vv1) < len(vv2) {
		if vv2[len(vv2)-1] == "0" {
			return 0
		}

		return -1
	}

	if len(vv1) == len(vv2) {
		return 0
	}

	if len(vv1) > len(vv2) {
		if vv1[len(vv1)-1] == "0" {
			return 0
		}

	}

	return 1
}

func CompareVersionII(version1, version2 string) int {
	m, n := len(version1), len(version2)

	i, j := 0, 0
	isDigit := func(c byte) bool {
		return c >= '0' && c <= '9'
	}

	var num1, num2 int
	for i < m || j < n {
		num1 = 0
		for i < m && isDigit(version1[i]) {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}

		i++

		num2 = 0
		for j < n && isDigit(version2[j]) {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}

		j++
		if num1 < num2 {
			return -1
		}

		if num1 > num2 {
			return 1
		}
	}

	return 0
}
