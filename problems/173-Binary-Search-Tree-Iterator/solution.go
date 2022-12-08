/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nodes []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var iter BSTIterator
	iter.nodes = make([]*TreeNode, 0)
	iter.FindNext(root)
	return iter
}

func (this *BSTIterator) Next() int {
	next := this.nodes[len(this.nodes)-1]
	this.nodes = this.nodes[:len(this.nodes)-1]
	if next.Right != nil {
		this.FindNext(next.Right)
	}
	return next.Val
}

func (this *BSTIterator) HasNext() bool {
	if len(this.nodes) == 0 {
		return false
	} else {
		return true
	}
}

func (this *BSTIterator) FindNext(root *TreeNode) {
	for root != nil {
		this.nodes = append(this.nodes, root)
		root = root.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

