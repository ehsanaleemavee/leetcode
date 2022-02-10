/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	first_pointer := head
	second_pointer := head

	for second_pointer != nil {

		if second_pointer.Next != nil {
			second_pointer = second_pointer.Next.Next
		}

		if second_pointer == nil || second_pointer.Next == nil {
			break
		}

		first_pointer = first_pointer.Next
	}

	mid := first_pointer
	if second_pointer != nil {
		// even nodes
		mid = first_pointer.Next
	}

	//reverse rest half
	reverseList := new(ListNode)
	reverseList.Val = mid.Val

	//mid = mid.Next
	for mid.Next != nil {
		reverseList = pushToTop(mid.Next.Val, reverseList)
		mid = mid.Next
	}

	for reverseList != nil {

		if reverseList.Val != head.Val {
			return false
		}
		reverseList = reverseList.Next
		head = head.Next
	}

	return true
}

func pushToTop(val int, head *ListNode) *ListNode {

	node := new(ListNode)
	node.Val = val
	node.Next = head

	return node

}