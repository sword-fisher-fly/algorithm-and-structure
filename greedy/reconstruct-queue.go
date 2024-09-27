package greedy

import (
	"sort"
)

// people[i]=(hi, ki)表示第i个人的身高为hi, 前面刚好有ki个身高大于等于hi的人
// Input: [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
// Sorted:  [[7,0], [7,1], [6,1], [5,0], [5,2], [4,4]]
// Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]

type People [][2]int

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	}

	return p[i][0] > p[j][0]
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ReconstructQueue(peoples [][2]int) [][2]int {
	res := [][2]int{}

	sort.Sort(People(peoples))

	res = append(res, peoples[0])

	for i := 1; i < len(peoples); i++ {
		startIndex := 0
		pos := peoples[i][1]
		for k := 0; k < pos; k++ {
			startIndex++
		}

		if startIndex == 0 {
			res = append([][2]int{peoples[i]}, res...)
		} else if startIndex == len(res) {
			res = append(res, peoples[i])
		} else {
			temp := make([][2]int, 0, len(res)+1)
			temp = append(temp, res[:startIndex]...)
			temp = append(temp, peoples[i])
			temp = append(temp, res[startIndex:]...)
			res = temp
		}
	}

	return res
}

func ReconstructQueueII(peoples [][2]int) [][2]int {
	res := make([][2]int, len(peoples))

	sort.Sort(People(peoples))

	res[0] = peoples[0]

	for i := 1; i < len(peoples); i++ {
		startIndex := 0
		pos := peoples[i][1]
		for k := 0; k < pos; k++ {
			startIndex++
		}

		if startIndex == i {
			res[startIndex] = peoples[i]
		} else {
			// startIndex...(len-2) -> startIndex+1 .. (len-1)
			// for k := len(res) - 1; k > startIndex; k-- {
			// 优化, 此时res填充完以后为i个有效元素
			for k := i; k > startIndex; k-- {

				res[k] = res[k-1]
			}

			res[startIndex] = peoples[i]
		}
	}

	return res
}
