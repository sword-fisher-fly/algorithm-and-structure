package str

import "strings"

// 1) +0        true
// 2) +021      false
// 3) +123      true
// 4) +1.       false
// 5) +1.01E+3  true
// 6) +1.01E+03 false
// 7) +1.01     true
// 8) +1.1.     false
// 9) +1.1E     false
// 10) +1.1E3   true
// 11) +1.1EE   false
// 12) +1.1E-3  true
// 13) +1.1E.   false
func IsValidScientific(s string) bool {
	hasSignE := false

	s = strings.TrimSpace(s)

	s = strings.ToLower(s)

	fields := strings.Split(s, "e")
	if len(fields) == 2 {
		hasSignE = true
	}

	if hasSignE {
		return isValidRealNumber(fields[0]) && isNonZeroInterger(fields[1])
	}

	return isValidRealNumber(s)
}

// func IsValidScientific(s string) bool {
// 	s = strings.TrimSpace(s) // 去除字符串首尾的空白字符

// 	hasDot := false           // 是否有小数点
// 	hasE := false             // 是否有 'e' 或 'E'
// 	hasSignAfterE := true     // 'e' 或 'E' 后是否有符号
// 	isNegativeAfterE := false // 'e' 或 'E' 后是否为负

// 	for i, r := range s {
// 		switch r {
// 		case '.', 'e', 'E':
// 			if r == '.' {
// 				if hasDot || hasE { // 如果已经有小数点或者'e'/'E'，则无效
// 					return false
// 				}
// 				hasDot = true
// 			} else if r == 'e' || r == 'E' {
// 				if hasE || len(s) == i+1 { // 如果已经有'e'/'E'或者这是最后一个字符，则无效
// 					return false
// 				}
// 				hasE = true
// 				hasSignAfterE = false
// 				isNegativeAfterE = false
// 			}
// 		case '+', '-':
// 			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' { // 只能在字符串开始或紧跟'e'/'E'后出现
// 				return false
// 			}
// 			if r == '-' {
// 				isNegativeAfterE = true
// 			}
// 			hasSignAfterE = true
// 		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
// 			if isNegativeAfterE && r == '0' && !hasSignAfterE { // 'e'/'E'后第一个非'+'/-'字符必须是非零数字
// 				return false
// 			}
// 			hasSignAfterE = false
// 		default:
// 			return false // 其他字符都是无效的
// 		}
// 	}

// 	// 如果没有数字，结果无效
// 	return len(s) > 0 && (strings.ContainsAny(s, "0123456789"))
// }
