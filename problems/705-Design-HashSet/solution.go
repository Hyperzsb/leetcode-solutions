type MyListNode struct {
	Val  int
	Prev *MyListNode
	Next *MyListNode
}

type MyHashSet struct {
	MaxBucket int
	Buckets   []*MyListNode
}

func Constructor() MyHashSet {
	mhs := MyHashSet{
		MaxBucket: 100,
		Buckets:   make([]*MyListNode, 100),
	}

	for i := range mhs.Buckets {
		mhs.Buckets[i] = &MyListNode{
			Val:  -1,
			Prev: nil,
			Next: nil,
		}
	}

	return mhs
}

func (this *MyHashSet) Add(key int) {
	idx := key % this.MaxBucket

	head := this.Buckets[idx]
	for head.Next != nil {
		head = head.Next
		if head.Val == key {
			return
		}
	}

	head.Next = &MyListNode{
		Val:  key,
		Prev: head,
		Next: nil,
	}
}

func (this *MyHashSet) Remove(key int) {
	idx := key % this.MaxBucket

	head := this.Buckets[idx]
	if head.Next == nil {
		return
	}

	flag := false
	for head.Next != nil {
		head = head.Next
		if head.Val == key {
			flag = true
			break
		}
	}

	if flag {
		head.Prev.Next = head.Next
		if head.Next != nil {
			head.Next.Prev = head.Prev
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	idx := key % this.MaxBucket

	head := this.Buckets[idx]
	for head.Next != nil {
		head = head.Next
		if head.Val == key {
			return true
		}
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

