type MListNode struct {
    Key int
    Val int
    Prev *MListNode
    Next *MListNode
}

type LRUCache struct {
    Capacity int
    Size int
    NodeMap map[int]*MListNode
    Head *MListNode
    Tail *MListNode
}


func Constructor(capacity int) LRUCache {
    lruCache := LRUCache{
        Capacity: capacity,
        Size: 0,
        NodeMap: make(map[int]*MListNode),
        Head: &MListNode{},
        Tail: &MListNode{},
    }

    lruCache.Head.Next = lruCache.Tail
    lruCache.Head.Prev = nil
    lruCache.Tail.Next = nil
    lruCache.Tail.Prev = lruCache.Head

    return lruCache
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.NodeMap[key]; ok {
        node.Next.Prev = node.Prev
        node.Prev.Next = node.Next

        node.Next = this.Head.Next
        this.Head.Next.Prev = node
        node.Prev = this.Head
        this.Head.Next = node

        return node.Val
    } else {
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.NodeMap[key]; ok {
        node.Val = value

        node.Next.Prev = node.Prev
        node.Prev.Next = node.Next

        node.Next = this.Head.Next
        this.Head.Next.Prev = node
        node.Prev = this.Head
        this.Head.Next = node

        return
    }

    if this.Size < this.Capacity {
        newNode := &MListNode{
            Key: key,
            Val: value,
        }

        this.NodeMap[key] = newNode

        newNode.Next = this.Head.Next
        this.Head.Next.Prev = newNode
        newNode.Prev = this.Head
        this.Head.Next = newNode

        this.Size++

        return
    }

    oldNode := this.Tail.Prev
    delete(this.NodeMap, oldNode.Key)

    oldNode.Prev.Next = this.Tail
    this.Tail.Prev = oldNode.Prev

    newNode := &MListNode{
        Key: key,
        Val: value,
    }

    this.NodeMap[key] = newNode

    newNode.Next = this.Head.Next
    this.Head.Next.Prev = newNode
    newNode.Prev = this.Head
    this.Head.Next = newNode
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

