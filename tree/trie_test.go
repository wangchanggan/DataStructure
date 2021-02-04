package tree

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("hi")
	fmt.Println(trie.Find("hi"))
}
