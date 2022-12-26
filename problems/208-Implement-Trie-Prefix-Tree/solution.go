type Node struct {
	Val      byte
	End      bool
	Children []*Node
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	root := Node{0, true, make([]*Node, 26)}
	return Trie{&root}
}

func (this *Trie) Insert(word string) {
	current := this.root
	for i := 0; i < len(word); i++ {
		if current.Children[word[i]-'a'] == nil {
			newNode := Node{word[i] - 'a', false, make([]*Node, 26)}
			current.Children[word[i]-'a'] = &newNode
		}
		current = current.Children[word[i]-'a']
	}
	current.End = true
}

func (this *Trie) Search(word string) bool {
	current := this.root
	for i := 0; i < len(word); i++ {
		if current.Children[word[i]-'a'] == nil {
			return false
		}
		current = current.Children[word[i]-'a']
	}
	return current.End
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root
	for i := 0; i < len(prefix); i++ {
		if current.Children[prefix[i]-'a'] == nil {
			return false
		}
		current = current.Children[prefix[i]-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

