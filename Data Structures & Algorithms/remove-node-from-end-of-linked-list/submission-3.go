/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
    slow, fast := dummy, head

	for fast != nil && n > 0 {
		fast = fast.Next
		n--
	}

	// Maybe here fast is already nil, the list is not big enough for the n element
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Slow is the node right before the eviction : 
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}
