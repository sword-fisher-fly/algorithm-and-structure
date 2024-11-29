package special

import "sort"

type persons [][]int

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	if p[i][0] > p[j][0] {
		return true
	}

	return p[i][0] == p[j][0] && p[i][1] < p[j][1]
}

func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ReconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))
	sort.Sort(persons(people))

	res[0] = people[0]

	for i := 1; i < len(people); i++ {
		idx := people[i][1]
		// copy(res[idx+1:], res[
		if idx < i {
			for j := len(res) - 1; j > idx; j-- {
				res[j] = res[j-1]
			}
		}
		res[idx] = people[i]
	}

	return res
}
