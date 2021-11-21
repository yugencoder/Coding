package trie

type Node struct {
	Val   rune
	IsEnd bool
	Next  map[rune]*Node
}

type Trie struct {
	root *Node
}

func New() Trie {
	return Trie{
		root: &Node{
			Next: make(map[rune]*Node),
		},
	}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		panic("Length of word must > 0.")
	}

	var (
		runes   = []rune(word)
		curNode = t.root
		tmp     *Node
		ok      bool
	)

	for _, letter := range runes {
		if tmp, ok = curNode.Next[letter]; !ok {
			tmp = &Node{Val: letter, Next: make(map[rune]*Node)}
			curNode.Next[letter] = tmp
		}
		curNode = tmp
	}
	curNode.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	var (
		runes   = []rune(word)
		curNode = t.root
		ok      bool
	)

	for _, letter := range runes {
		if curNode, ok = curNode.Next[letter]; !ok {
			return false
		}
	}

	if !curNode.IsEnd {
		return false
	}

	return true
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	var (
		runes   = []rune(prefix)
		curNode = t.root
		ok      bool
	)

	for _, letter := range runes {
		if curNode, ok = curNode.Next[letter]; !ok {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := New();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
