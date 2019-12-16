package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func primarySchoolAddition(l1 *ListNode, l2 *ListNode) *ListNode {
	var list ListNode = ListNode{}
	current := &list
	hold := 0
	for l1 != nil && l2 != nil {
		add := l1.Val + l2.Val + hold
		if add >= 10 {
			add = add % 10
			hold = 1
		} else {
			hold = 0
		}
		current.Val = add
		if l1.Next != nil || l2.Next != nil || hold == 1 {
			current.Next = &ListNode{}
		}
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		add := l1.Val + hold
		if add >= 10 {
			add = add % 10
			hold = 1
		} else {
			hold = 0
		}
		current.Val = add
		if l1.Next != nil || hold == 1 {
			current.Next = &ListNode{}
		}
		current = current.Next
		l1 = l1.Next
	}

	for l2 != nil {
		add := l2.Val + hold
		if add >= 10 {
			add = add % 10
			hold = 1
		} else {
			hold = 0
		}
		current.Val = add
		if l2.Next != nil || hold == 1 {
			current.Next = &ListNode{}
		}
		current = current.Next
		l2 = l2.Next
	}
	if hold == 1 {
		current.Val = hold
	} else {
		current = nil
	}

	return &list
}

//you can't convert to int andd do the addition because you may overflow
//so do the addition like you learned it in primary school
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return primarySchoolAddition(l1, l2)
}
