package interview

// 单词多会超时
func FindWords(board [][]byte, words []string) []string {
	directions := [4][2]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	clean := func() {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				visited[i][j] = false
			}
		}
	}

	var dfs func(int, int, int, string) bool
	dfs = func(x, y int, startIdx int, word string) bool {
		if startIdx == len(word)-1 {
			return word[startIdx] == board[x][y]
		}

		if word[startIdx] == board[x][y] {
			visited[x][y] = true
			for _, d := range directions {
				nextX, nextY := x+d[0], y+d[1]
				if nextX >= 0 && nextX < len(board) && nextY >= 0 && nextY < len(board[0]) && !visited[nextX][nextY] {
					if dfs(nextX, nextY, startIdx+1, word) {
						return true
					}
				}
			}
			visited[x][y] = false
		}

		return false
	}

	res := []string{}
	for _, w := range words {
		clean()
	LOOP:
		// fmt.Printf("word=%s\n", w)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dfs(i, j, 0, w) {
					res = append(res, w)
					break LOOP
				}
			}
		}
	}

	return res
}

type Trie struct {
	children map[byte]*Trie
	word     string
}

func (t *Trie) Insert(word string) {
	node := t
	for i := range word {
		ch := word[i]
		if node.children[ch] == nil {
			node.children[ch] = &Trie{children: map[byte]*Trie{}}
		}
		node = node.children[ch]
	}
	node.word = word
}

func FindWordsII(board [][]byte, words []string) (ans []string) {
	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	t := &Trie{children: map[byte]*Trie{}}
	for _, word := range words {
		t.Insert(word)
	}

	m, n := len(board), len(board[0])

	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		nxt := node.children[ch]
		if nxt == nil {
			return
		}

		if nxt.word != "" {
			ans = append(ans, nxt.word)
			nxt.word = ""
		}

		if len(nxt.children) > 0 {
			board[x][y] = '#'
			for _, d := range dirs {
				nx, ny := x+d.x, y+d.y
				if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
					dfs(nxt, nx, ny)
				}
			}
			board[x][y] = ch
		}

		if len(nxt.children) == 0 {
			delete(node.children, ch)
		}
	}

	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}

	return
}
