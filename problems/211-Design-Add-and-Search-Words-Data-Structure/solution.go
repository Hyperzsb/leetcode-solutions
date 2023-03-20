type TNode struct {
	Val      byte
	End      bool
	Children map[byte]*TNode
}

type WordDictionary struct {
	Root *TNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		Root: &TNode{
			Val:      0,
			End:      true,
			Children: make(map[byte]*TNode),
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	current := this.Root
	for i := 0; i < len(word); i++ {
		if current.Children[word[i]] == nil {
			current.Children[word[i]] = &TNode{
				Val:      word[i],
				End:      false,
				Children: make(map[byte]*TNode),
			}
		}

		current = current.Children[word[i]]
	}
	current.End = true
}

func searchTree(word string, root *TNode) bool {
	if len(word) == 0 {
		return true
	}

	if len(word) == 1 {
		if word[0] == '.' {
			for _, child := range root.Children {
				if child.End {
					return true
				}
			}

			return false
		} else {
			if root.Children[word[0]] != nil && root.Children[word[0]].End {
				return true
			}

			return false
		}
	}

	if word[0] == '.' {
		for _, child := range root.Children {
			if searchTree(string(word[1:]), child) {
				return true
			}
		}

		return false
	}

	if root.Children[word[0]] == nil {
		return false
	}

	return searchTree(string(word[1:]), root.Children[word[0]])
}

func (this *WordDictionary) Search(word string) bool {
	return searchTree(word, this.Root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

