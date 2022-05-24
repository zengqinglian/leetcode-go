package swappingNodesinaLinkedList

import "testing"

func TestFunc(t *testing.T) {
	head := ListNode{Val: 7}
	n2 := ListNode{Val: 9}
	// n3 := ListNode{Val: 6}
	// n4 := ListNode{Val: 6}
	// n5 := ListNode{Val: 7}
	// n6 := ListNode{Val: 8}
	// n7 := ListNode{Val: 3}
	// n8 := ListNode{Val: 0}
	// n9 := ListNode{Val: 9}
	// n10 := ListNode{Val: 5}
	head.Next = &n2
	// n2.Next = &n3
	// n3.Next = &n4
	// n4.Next = &n5
	// n5.Next = &n6
	// n6.Next = &n7
	// n7.Next = &n8
	// n8.Next = &n9
	// n9.Next = &n10
	swapNodes(&head, 1)
}
