package greedy

import "sort"

// 每瓶汽水5美元，每人购买一份，每人身上的钱为[5,10,20]面值
// 给定一个数组表示每人身上的钱，判断是否能每人买一瓶汽水并找零
func lemonadeChange(bills []int) bool {
	sort.Ints(bills)

	five, ten, twenty := 0, 0, 0

	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			if five < 1 {
				return false
			}
			five--
			ten++
		case 20:
			if five >= 1 && ten >= 1 {
				five--
				ten--
				twenty++
			} else if five >= 3 {
				five -= 3
				twenty++
			} else {
				return false
			}
		}
	}

	return true
}
