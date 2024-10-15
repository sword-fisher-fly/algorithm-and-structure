package str

// 某人在键盘输入打字。 偶尔，在输入字符c时会长按，而字符可能输入1次或多次。
// 检查键盘输入的字符typed。如果它可能是某个文本的输入，则返回True。

// 输入：name = "alex", typed = "aaleex"
// 输出：true
// 思路:
//  1. i, j分别表示name和typed的下标索引
//  2. 如果name[i] == typed[j], 则i++, j++
//  3. 如果name[i]!= typed[j],
//     a. 如果j == 0, 则返回false
//     b. 如果typed[j] == typed[j-1], 则j++
//     c. 如果name[i] != typed[j], 则返回false, 否则继续比较
//  4. 如果i < len(name), 则返回false
//  5. 如果j < len(typed), 则判断typed[j] == typed[j-1], 如果是则j++, 否则返回false
func IsLongPressName(name string, typed string) bool {
	i, j := 0, 0

	for i < len(name) && j < len(typed) {
		if name[i] == typed[j] {
			i++
			j++
			continue
		}

		if j == 0 {
			return false
		}

		for j < len(typed) && typed[j] == typed[j-1] {
			j++
		}

		// j 可能越界, 所以需要判断j是否越界
		if j == len(typed) || name[i] != typed[j] {
			return false
		}
	}

	if i < len(name) {
		return false
	}

	// if j < len(typed) {
	// 	for j < len(typed) && typed[j] != typed[j-1] {
	// 		return false
	// 	}

	// 	j++
	// }

	for j < len(typed) {
		if typed[j] != typed[j-1] {
			return false
		}

		j++
	}

	return true
}
