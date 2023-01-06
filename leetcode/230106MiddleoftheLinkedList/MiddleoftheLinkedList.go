/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	add := head
	count := 1
	for add.Next != nil {
		count++
		add = add.Next
	}
	count = count/2 + 1
	add = head
	for count > 1 {
		count--
		add = add.Next
	}

	return add
}