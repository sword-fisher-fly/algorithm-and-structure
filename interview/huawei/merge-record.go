package huawei

import (
	"fmt"
	"sort"
)

// https://www.nowcoder.com/practice/de044e89123f4a7482bd2b214a685201?tpId=37&tqId=21231&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D37&difficulty=undefined&judgeStatus=undefined&tags=&title=
type pair struct {
	index int
	value int
}

type ps []pair

func (p ps) Len() int {
	return len(p)
}

func (p ps) Less(i, j int) bool {
	return p[i].index < p[j].index
}

func (p ps) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (r pair) String() string {
	return fmt.Sprintf("%d %d", r.index, r.value)
}

func MergeRecord(pairs [][2]int) [][2]int {
	var rs ps

	for i := 0; i < len(pairs); i++ {
		rs = append(rs, pair{index: pairs[i][0], value: pairs[i][1]})
	}

	res := [][2]int{}
	sort.Sort(rs)

	mergeRecord := func(s []pair) pair {
		for i := 1; i < len(s); i++ {
			s[0].value += s[i].value
		}

		return s[0]
	}

	i, j := 0, 0
	l := len(rs)
	for i < l {
		j = i + 1
		for j < l && rs[i].index == rs[j].index {
			j++
		}
		m := mergeRecord(rs[i:j])
		res = append(res, [2]int{m.index, m.value})
		i = j
	}

	return res
}
