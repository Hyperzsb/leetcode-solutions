/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	listMap := make(map[int]*ListNode)
	for i := 0; i < len(lists); i++ {
		listMap[i] = lists[i]
	}

	var head *ListNode
	for key, list := range listMap {
		if list == nil {
			delete(listMap, key)
			continue
		}

		if head == nil {
			head = list
			continue
		}

		if head.Val > list.Val {
			head = list
		}
	}

	previous, current := head, (*ListNode)(nil)
	for previous != nil && len(listMap) > 0 {
		current = nil
		for key, list := range listMap {
			if list == nil {
				delete(listMap, key)
				continue
			}

			if list == previous {
				listMap[key] = list.Next
				list = list.Next
			}

			if list == nil {
				delete(listMap, key)
				continue
			}

			if current == nil {
				current = list
				continue
			}

			if current.Val > list.Val {
				current = list
			}
		}

		previous.Next = current
		previous = current
	}

	return head
}

