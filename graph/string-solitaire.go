package graph

func StringSolitaire(beginWord, endWord string, words []string) int {
	wordDict := make(map[string]bool)
	for i := range words {
		wordDict[words[i]] = true
	}

	// 记录到该单词长度
	visitedMap := make(map[string]int)
	visitedMap[beginWord] = 1
	queue := []string{beginWord}

	for len(queue) > 0 {
		curWord := queue[0]
		queue = queue[1:]
		path := visitedMap[curWord]

		for i := 0; i < len(curWord); i++ {
			newWord := []byte(curWord)
			for j := 0; j < 26; j++ {
				newWord[i] = byte(j + 'a')
				strNewWord := string(newWord)
				if strNewWord == endWord {
					return path + 1
				}

				if wordDict[strNewWord] {
					if _, ok := visitedMap[strNewWord]; !ok {
						visitedMap[strNewWord] = path + 1
						queue = append(queue, strNewWord)
					}
				}
			}
		}
	}

	return 0
}
