package huawei

func ParseArguments(input string) []string {
	res := []string{}

	start := 0
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			res = append(res, input[start:i])
			start = i + 1
		}

		if i == len(input)-1 {
			res = append(res, input[start:i+1])
		}

		if input[i] == '"' {
			start = i + 1
			j := i + 1
			for j < len(input) {
				if input[j] == '"' {
					res = append(res, input[start:j])
					break
				}
				j++
			}

			i = j + 1     // 空格, 再经过for i++ 指向下一个字符
			start = j + 2 // 空格后第一个字符
		}

	}

	return res
}
