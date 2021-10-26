package intersectionoftwolinkedlists

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB ==nil {
        return nil
    }
    a := headA
    b := headB
    
    for a!=b {
        if a == nil {
            a = headB
        } else {
			a = a.Next
        }
        if b == nil {
            b = headA
            
        }else {
            b = b.Next
        }
        
    }
    return a
}