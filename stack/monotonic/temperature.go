package monotonic

// [73, 74, 75, 71, 69, 72, 76, 73]
// 输出数组: 要想观测到更高的气温，至少需要等待的天数
func DailyTemperaturesByForce(temperatures []int) []int {
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures)-1; i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}

	return result
}
func DailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	if len(temperatures) <= 1 {
		return res
	}

	stack := []int{0} // 存储元素的index

	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
				res[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, i)
		}
	}

	return res
}

func DailyTemperaturesII(temperatures []int) []int {
	res := make([]int, len(temperatures))

	if len(temperatures) <= 1 {
		return res
	}

	stack := []int{0} // 存储温度的下标, 代表第i天

	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
