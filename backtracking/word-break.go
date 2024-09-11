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
