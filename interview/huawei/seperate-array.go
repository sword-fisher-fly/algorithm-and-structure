package huawei

//https://www.nowcoder.com/practice/9af744a3517440508dbeb297020aca86?tpId=37&tqId=21316&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3Fpage%3D2%26tpId%3D37%26type%3D37&difficulty=undefined&judgeStatus=undefined&tags=&title=
func SeperateArray(arr []int) bool {
	total := len(arr)
	nums := []int{}
	sum, sum3, sum5 := 0, 0, 0
	for i := 0; i < total; i++ {
		cur := nums[i]
		if cur%5 == 0 {
			sum5 += cur
		} else if cur%3 == 0 {
			sum3 += cur
		} else {
			nums = append(nums, cur)
		}
		sum += cur
	}

	target := sum/2 - sum3
	if sum%2 != 0 {
		return false
	}

	var dfs func(int, []int, int) bool
	dfs = func(startIndex int, arr []int, target int) bool {
		if startIndex == len(arr) {
			return target == 0
		}

		return dfs(startIndex+1, arr, target-arr[startIndex]) || dfs(startIndex+1, arr, target)
	}

	return dfs(0, arr, target)
}

func SeperateArrayII(arr []int) bool {
	threeArr, fiveArr, otherArr := []int{}, []int{}, []int{}
	n := len(arr)
	for i := 0; i < n; i++ {
		cur := arr[i]
		if cur%5 == 0 {
			fiveArr = append(fiveArr, cur)
		} else if cur%3 == 0 {
			threeArr = append(threeArr, cur)
		} else {
			otherArr = append(otherArr, cur)
		}
	}

	sum := func(arr []int) int {
		s := 0
		for i := range arr {
			s += arr[i]
		}
		return s
	}

	var dfs func([]int, []int, []int) bool

	dfs = func(three, five, other []int) bool {
		if len(other) == 0 {
			if sum(three) == sum(five) {
				return true
			}

			return false
		}

		if dfs(append(three, other[:1]...), five, other[1:]) {
			return true
		}

		if dfs(three, append(five, other[:1]...), other[1:]) {
			return true
		}

		return false
	}

	return dfs(threeArr, fiveArr, otherArr)
}
