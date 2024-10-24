package offer

import "strings"

// TODO:
// 1) 哈希表
// 2）字典树
func ReplaceWord(s string, dictionary []string) string {
	dictSet := make(map[string]bool)
	for i := range dictionary {
		dictSet[dictionary[i]] = true
	}

	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		for j := 1; j <= len(words[i]); j++ {
			if _, ok := dictSet[words[i][:j]]; ok {
				words[i] = words[i][:j]
				break
			}
		}
	}

	return strings.Join(words, " ")
}

func ReplaceWordII(s string, dictionary []string) string {
	type trie map[rune]trie

	root := trie{}
	for _, w := range dictionary {
		cur := root
		for _, c := range w {
			if cur[c] == nil {
				cur[c] = trie{}
			}
			cur = cur[c]
		}
		cur['#'] = trie{}
	}

	words := strings.Fields(s)
	for i, w := range words {
		cur := root
		for j, c := range w {
			if cur['#'] != nil {
				words[i] = w[:j]
				break
			}
			if cur[c] == nil {
				break
			}
			cur = cur[c]
		}
	}

	return strings.Join(words, " ")
}
