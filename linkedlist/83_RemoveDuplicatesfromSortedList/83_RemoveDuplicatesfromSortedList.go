/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	temp := head //	var temp1 *ListNode = head

	m := make(map[int]*ListNode)

	for temp != nil {
		_, present := m[temp.Val]

		if present == false {
			m[temp.Val] = temp
		}

		temp = temp.Next
	}

	s := make([]int, 0, len(m))

	for key := range m {
		s = append(s, key)
	}

	sort.Ints(s)

	var list *ListNode
	var node *ListNode
	for key := range s {

		if list == nil {
			node = new(ListNode)
			node.Val = s[key]
			list = node
			continue
		}

		node.Next = new(ListNode)
		node = node.Next
		node.Val = s[key]
	}

	return list
}