package graph

import "fmt"

// BFS遍历如何记录路径??
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
				// 优化
				if curWord == strNewWord {
					continue
				}

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

func StringSolitaireByDFS(beginWord, endWord string, words []string) int {
	wordDict := make(map[string]bool)
	for i := range words {
		wordDict[words[i]] = true
	}

	// 记录到该单词长度
	visitedMap := make(map[string]struct{})
	visitedMap[beginWord] = struct{}{}

	traversalPath := []string{beginWord}
	minPath := []string{}
	resPath := [][]string{}

	minLen := 0

	var dfs func(string)

	dfs = func(curWord string) {
		for i := 0; i < len(curWord); i++ {
			newWord := []byte(curWord)
			for j := 0; j < 26; j++ {
				newWord[i] = byte(j + 'a')
				strNewWord := string(newWord)
				// 优化
				if curWord == strNewWord {
					continue
				}

				if strNewWord == endWord {
					temp := make([]string, len(traversalPath))
					copy(temp, traversalPath)
					temp = append(temp, endWord)

					if len(temp) < minLen || minLen == 0 {
						minPath = make([]string, len(temp))
						copy(minPath, temp)
						minLen = len(minPath)
					}

					resPath = append(resPath, temp)
					return
				}

				if wordDict[strNewWord] {
					if _, ok := visitedMap[strNewWord]; !ok {
						visitedMap[strNewWord] = struct{}{}
						traversalPath = append(traversalPath, strNewWord)
						dfs(strNewWord)
						traversalPath = traversalPath[:len(traversalPath)-1]
						delete(visitedMap, strNewWord)
					}
				}
			}
		}
	}

	dfs(beginWord)

	fmt.Printf("minPath: %v\n", minPath)

	fmt.Println("----   All traversal path   ----")
	for i, r := range resPath {
		fmt.Printf("Path %d: %v\n", i+1, r)
	}

	return minLen
}
