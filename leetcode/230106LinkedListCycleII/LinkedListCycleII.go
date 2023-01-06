//https://leetcode.com/problems/linked-list-cycle-ii/description/?envType=study-plan&id=level-1
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil && m[head] != true {
		m[head] = true
		head = head.Next
	}
	return head
}