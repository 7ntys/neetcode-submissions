/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var tail *ListNode = nil
	for l1 != nil && l2 != nil {
		add := l1.Val + l2.Val + carry
		carry = 0
		if add > 9 {
			carry = add / 10
			add = add % 10
		}

		node := &ListNode{Val: add, Next: tail}
		tail = node

		l1 = l1.Next
		l2 = l2.Next
	}
	
	for l1 != nil {
		add := l1.Val + carry 
		carry = 0
		if add > 9 {
			carry = add / 10
			add = add % 10
		}

		node := &ListNode{Val: add, Next: tail}
		tail = node

		l1 = l1.Next
	}

	for l2 != nil {
		add := l2.Val + carry 
		carry = 0
		if add > 9 {
			carry = add / 10
			add = add % 10
		}

		node := &ListNode{Val: add, Next: tail}
		tail = node

		l2 = l2.Next
	}

	if carry > 0 {
		node := &ListNode{Val: carry, Next: tail}
		tail = node
	}

	return reverseLL(tail)
}

func reverseLL(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}