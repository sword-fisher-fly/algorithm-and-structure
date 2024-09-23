package trie

import "testing"

func Test_TrieTree(t *testing.T) {
	tr := NewTrieTree()

	tr.Insert("abc")
	tr.Insert("work")
	tr.Insert("world")

	t.Logf("Search('abc')=%v", tr.Search("abc"))
	t.Logf("Search('word')=%v", tr.Search("word"))
	t.Logf("Search('work')=%v", tr.Search("work"))
	t.Logf("Search('worlk')=%v", tr.Search("worlk"))
	t.Logf("Search('world')=%v", tr.Search("world"))

}
