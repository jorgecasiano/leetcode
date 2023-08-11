/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	var acc = 0

	for l1 != nil || l2 != nil {
		val1 := getValOrZero(l1)
		val2 := getValOrZero(l2)
		sum := val1 + val2 + acc

		if sum < 10 {
			acc = 0
		} else {
			sum = sum - 10
			acc = 1
		}

		newNode := &ListNode{Val: sum}
		current.Next = newNode
		current = newNode

		l1 = getNext(l1)
		l2 = getNext(l2)
	}

	if acc == 1 {
		newNode := &ListNode{Val: 1}
		current.Next = newNode
	}

	return result.Next
}

func getValOrZero(l *ListNode) int {
	if (l != nil) {
		return l.Val
	}

	return 0
}

func getNext(l *ListNode) *ListNode {
	if (l != nil) {
		return l.Next
	}

	return nil
}