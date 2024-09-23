package trie

type TrieNode struct {
	children [26]*TrieNode
	flag     bool
}
type TrieTree struct {
	root *TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{root: &TrieNode{}}
}

func (t *TrieTree) Insert(s string) {
	node := t.root
	for i := range s {
		idx := s[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}

		node = node.children[idx]
	}

	node.flag = true
}

func (t *TrieTree) Search(s string) bool {
	node := t.root

	for i := range s {
		idx := s[i] - 'a'
		if node.children[idx] == nil {
			return false
		}

		node = node.children[idx]
	}

	return node.flag
}
