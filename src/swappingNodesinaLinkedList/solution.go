package swappingNodesinaLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 208 ms, faster than 59.25% of Go online submissions for Swapping Nodes in a Linked List.
Memory Usage: 9.7 MB, less than 77.64% of Go online submissions for Swapping Nodes in a Linked List.
*/
func swapNodes(head *ListNode, k int) *ListNode {
	var pre *ListNode
	var cur *ListNode
	var targetKPre *ListNode
	var targetK *ListNode
	cur = head

	cnt := 0
	for cur != nil {
		if cnt == k-1 {
			targetKPre = pre
			targetK = cur
		}
		pre = cur
		cur = cur.Next
		cnt++
	}

	if cnt == 1 {
		return head
	}

	if (cnt/2) == (k-1) && cnt%2 == 1 {
		return head
	}

	cnt = cnt - k
	var targetK1Pre *ListNode
	var targetK1 *ListNode
	pre = nil
	cur = head

	for i := 0; i < cnt; i++ {
		pre = cur
		cur = cur.Next
	}
	targetK1Pre = pre
	targetK1 = cur

	originalKNext := targetK.Next
	originalK1Next := targetK1.Next

	if targetK.Next == targetK1 {
		targetK1.Next = targetK
		targetK.Next = originalK1Next
		if targetKPre != nil {
			targetKPre.Next = targetK1
		} else {
			head = targetK1
		}
	} else if targetK1.Next == targetK {
		targetK.Next = targetK1
		targetK1.Next = originalKNext
		if targetK1Pre != nil {
			targetK1Pre.Next = targetK
		} else {
			head = targetK
		}
	} else {

		if targetKPre != nil {
			targetKPre.Next = targetK1
		} else {
			head = targetK1
		}
		if targetK.Next == targetK1 {
			targetK1.Next = targetK
		} else {
			targetK1.Next = originalKNext
		}

		if targetK1Pre != nil {
			targetK1Pre.Next = targetK
		} else {
			head = targetK
		}
		if targetK1.Next == targetK {
			targetK.Next = targetK1
		} else {
			targetK.Next = originalK1Next
		}
	}

	return head
}
