package trie

// TODO:
// https://leetcode.cn/problems/short-encoding-of-words/

func MinimumLengthEncoding(words []string) int {
	return 0
}

func MinimumLengthEncodingII(words []string) int {
	wordDict := make(map[string]struct{})

	for _, word := range words {
		wordDict[word] = struct{}{}
	}

	for word := range wordDict {
		for i := 1; i < len(word); i++ {
			delete(wordDict, word[i:])
		}
	}

	ans := 0
	for word := range wordDict {
		ans += len(word) + 1
	}

	return ans
}

type Trie struct {
	children [26]*Trie
}

func MinimumLengthEncodingByTrie(words []string) int {
	root := &Trie{}
	for _, w := range words {
		cur := root
		for i := len(w) - 1; i >= 0; i-- {
			idx := w[i] - 'a'
			if cur.children[idx] == nil {
				cur.children[idx] = &Trie{}
			}

			cur = cur.children[idx]
		}
	}

	var dfs func(*Trie, int) int
	dfs = func(t *Trie, l int) int {
		isLeaf, ans := true, 0

		for i := 0; i < 26; i++ {
			if t.children[i] != nil {
				isLeaf = false
				ans += dfs(t.children[i], l+1)
			}
		}

		if isLeaf {
			ans += l
		}

		return ans
	}

	return dfs(root, 1)
}
