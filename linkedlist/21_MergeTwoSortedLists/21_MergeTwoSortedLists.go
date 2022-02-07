/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var head *ListNode
	var node *ListNode

	for list1 != nil || list2 != nil {

		data1 := 101
		data2 := 101
		pick := 101

		if list1 != nil {
			data1 = list1.Val
		}

		if list2 != nil {
			data2 = list2.Val
		}

		if data1 > data2 {
			pick = data2
			list2 = list2.Next
		} else {
			pick = data1
			list1 = list1.Next
		}

		if head == nil {
			node = new(ListNode)
			node.Val = pick
			head = node
			continue
		}

		node.Next = new(ListNode)
		node = node.Next
		node.Val = pick

	}

	if head == nil {
		return list1
	}

	return head
}