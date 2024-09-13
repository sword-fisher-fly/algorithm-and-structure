package backtracking

func workBreakHelper(s string, startIndex int, wordDict map[string]struct{}, memo map[int]bool) bool {
	if startIndex == len(s) {
		return true
	}

	if !memo[startIndex] {
		return false
	}

	for i := startIndex; i < len(s); i++ {
		subStr := s[startIndex : i+1]
		if _, ok := wordDict[subStr]; !ok {
			continue
		}
		return workBreakHelper(s, i+1, wordDict, memo)
	}

	memo[startIndex] = false

	return false
}

func WorkBreak(s string, wordDict []string) bool {
	memo := make(map[int]bool)
	for i := range s {
		memo[i] = true
	}

	wordMap := make(map[string]struct{}, len(wordDict))
	for i := range wordDict {
		wordMap[wordDict[i]] = struct{}{}
	}

	return workBreakHelper(s, 0, wordMap, memo)
}

func WordBreakII(s string, words []string) bool {
	if len(s) == 0 || len(words) == 0 {
		return false
	}

	wordDict := make(map[string]struct{}, len(words))
	for i := range words {
		wordDict[words[i]] = struct{}{}
	}

	// 优化: 初始化为true, 表示s[0:i]可分割
	memo := make([]bool, len(s)+1)
	for i := range memo {
		memo[i] = true
	}

	var backtracking func(s string, startIndex int, wordDict map[string]struct{}) bool
	backtracking = func(s string, startIndex int, wordDict map[string]struct{}) bool {
		if startIndex == len(s) {
			return true
		}

		if !memo[startIndex] {
			return false
		}

		for i := startIndex; i < len(s); i++ {
			subStr := s[startIndex : i+1]
			if _, ok := wordDict[subStr]; !ok {
				continue
			}

			return backtracking(s, i+1, wordDict)
		}

		memo[startIndex] = false

		return false
	}

	return backtracking(s, 0, wordDict)
}
