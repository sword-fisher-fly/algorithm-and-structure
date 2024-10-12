package offer

// 外星语言使用小写字母, 可能字母order的顺序不同
// 给定一组外星语书写的单词words, 以及其字母表的顺序order
// 当给定单词按照外星语字母顺序排序时, 返回true, 否则返回false

func IsAlienSorted(words []string, order string) bool {
	orderIndex := [26]int{}
	for i := range order {
		orderIndex[order[i]-'a'] = i
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		l1, l2 := len(words[i]), len(words[i+1])
		maxLen := max(l1, l2)
		for j := 0; j < maxLen; j++ {
			i1, i2 := -1, -1
			if j < l1 {
				i1 = orderIndex[w1[j]-'a']
			}

			if j < l2 {
				i2 = orderIndex[w2[j]-'a']
			}

			if i1 > i2 {
				return false
			}

			if i1 < i2 {
				break
			}
		}
	}

	return true
}
