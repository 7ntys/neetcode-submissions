/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Val:0, Next: head}
	prevGroup := dummy
	for head != nil {
		kth := getKth(head,k)
		if kth == nil {break}
		nextGroup := kth.Next

		// inverse from prevGroup.Next till kth :
		prev, curr := nextGroup, head
		for curr != nextGroup {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		prevGroup.Next = kth
		head.Next = nextGroup
		prevGroup = head
		head = nextGroup
	}
	return dummy.Next
}

func getKth(node *ListNode, k int) *ListNode {
	for i := 1; i < k && node != nil; i++ {
		node = node.Next
	}
	return node
}
