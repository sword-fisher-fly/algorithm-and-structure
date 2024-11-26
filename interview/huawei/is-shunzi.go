package huawei

import "sort"

//https://www.nowcoder.com/practice/762836f4d43d43ca9deb273b3de8e1f4?tpId=196&tqId=37092&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj&difficulty=undefined&judgeStatus=undefined&tags=&title=
func IsContinuous(numbers []int) bool {
	// write code here
	sort.Ints(numbers)
	countZero := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 0 {
			countZero++
		}
	}

	interrupt := 0
	for i := countZero + 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] {
			return false
		}

		interrupt += numbers[i] - numbers[i-1] - 1
	}

	if interrupt > countZero {
		return false
	}

	return true
}
