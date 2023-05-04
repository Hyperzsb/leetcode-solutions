type MyListNode struct {
	Val  byte
	Next *MyListNode
}

func predictPartyVictory(senate string) string {
	head := &MyListNode{senate[0], nil}
	tail := head
	for i := 1; i < len(senate); i++ {
		tail.Next = &MyListNode{senate[i], nil}
		tail = tail.Next
	}
	tail.Next = head

	for {
		senator, curr, next := head.Val, head, head.Next
		for next != head {
			if next.Val == senator {
				curr = curr.Next
				next = next.Next
			} else {
				curr.Next = next.Next
				break
			}
		}

		if next == head {
			break
		} else {
			head = head.Next
		}
	}

	result := ""
	if head.Val == 'R' {
		result = "Radiant"
	} else {
		result = "Dire"
	}

	return result
}

