type MListNode struct {
    Key int
    Val int
    Next *MListNode
}

type MyHashMap struct {
    Capacity int
    Buckets []*MListNode
}


func Constructor() MyHashMap {
    mhm := MyHashMap{
        Capacity: 1e4,
        Buckets: make([]*MListNode, 1e4),
    }

    for i := range mhm.Buckets {
        mhm.Buckets[i] = &MListNode{}
    }

    return mhm
}


func (this *MyHashMap) Put(key int, value int)  {
    hashKey := key % this.Capacity

    if this.Buckets[hashKey].Next == nil {
        this.Buckets[hashKey].Next = &MListNode{
            Key: key,
            Val: value,
        }
    } else {
        curr := this.Buckets[hashKey]

        exists := false
        for curr.Next != nil {
            if curr.Next.Key == key {
                curr.Next.Val = value
                exists = true
                break
            } else {
                curr = curr.Next
            }
        }

        if !exists {
            curr.Next = &MListNode {
                Key: key,
                Val: value,
            }
        }
    }
}


func (this *MyHashMap) Get(key int) int {
    hashKey := key % this.Capacity

    curr := this.Buckets[hashKey].Next
    for curr != nil {
        if curr.Key == key {
            break
        } else {
            curr = curr.Next
        }
    }

    if curr == nil {
        return -1
    } else {
        return curr.Val
    }
}


func (this *MyHashMap) Remove(key int)  {
    hashKey := key % this.Capacity

    prev, curr := this.Buckets[hashKey], this.Buckets[hashKey].Next
    for curr != nil {
        if curr.Key == key {
            break
        } else {
            prev = prev.Next
            curr = curr.Next
        }
    }

    if curr != nil {
        prev.Next = curr.Next
    } 
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

