package removeDuplicatesFromSortedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
        return nil
    }
    
    if head.Next == nil {
        return head
    }
    
    if head.Val == head.Next.Val {
        return deleteDuplicates(head.Next)
    }
    
    head.Next = deleteDuplicates(head.Next)
    return head
}

func deleteDuplicates01(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
