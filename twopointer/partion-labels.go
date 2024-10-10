package twopointer

// 字符串s由小写字母组成, 将字符串划分为尽可能多的片段, 同一个字母最多出现在一个片段中
// 求解: 表示每个字符串片段的长度列表
// 巧妙点: 在于不断更新右端的边界且保证当前扫描区间的字符都在其值的左边
func PartitionLabels(s string) []int {
	hash := [26]int{} // 记录每个字符在s中出现的最大索引

	for i := range s {
		hash[s[i]-'a'] = i
	}

	res := []int{}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		right = max(right, hash[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}

	return res
}
