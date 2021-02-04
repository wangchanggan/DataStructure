package tree

type TrieNode struct {
	Data         byte
	Children     [26]*TrieNode
	IsEndingByte bool
}

func NewTrieNode(data byte) *TrieNode {
	trieNode := new(TrieNode)
	trieNode.Data = data
	return trieNode
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	trie := new(Trie)
	trie.Root = NewTrieNode('/')
	return trie
}

func (trie *Trie) Insert(text string) {
	p := trie.Root
	for i := 0; i < len(text); i++ {
		index := text[i] - 'a'
		if p.Children[index] == nil {
			newTrieNode := NewTrieNode(text[i])
			p.Children[index] = newTrieNode
		}
		p = p.Children[index]
	}
	p.IsEndingByte = true
}

func (trie *Trie) Find(pattern string) bool {
	p := trie.Root
	for i := 0; i < len(pattern); i++ {
		index := pattern[i] - 'a'
		if p.Children[index] == nil {
			return false
		}
		p = p.Children[index]
	}

	if p.IsEndingByte == false {
		return false
	} else {
		return true
	}
}
