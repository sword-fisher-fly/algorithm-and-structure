package huawei

import (
	"sort"
	"strings"
)

//https://www.nowcoder.com/practice/c1f9561de1e240099bdb904765da9ad0?tpId=37&tqId=21325&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3Fpage%3D2%26tpId%3D37%26type%3D37&difficulty=undefined&judgeStatus=undefined&tags=&title=

type CharStatis struct {
	c     byte
	count int
}

type cs []*CharStatis

func (c cs) Len() int {
	return len(c)
}

func (c cs) Less(i, j int) bool {
	if c[i].count == c[j].count {
		return c[i].c < c[j].c
	}

	return c[i].count > c[j].count
}

func (c cs) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func CountAndSortChars(s string) string {
	cMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cMap[s[i]]++
	}

	var res []*CharStatis
	for c, s := range cMap {
		res = append(res, &CharStatis{c: c, count: s})
	}

	sort.Sort(cs(res))

	sb := strings.Builder{}
	for i := range res {
		sb.WriteByte(res[i].c)
	}

	return sb.String()
}
