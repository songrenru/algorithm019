/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil {
	// 	return head
	// }

	dummyHead := &ListNode{Next:head}
	pre := dummyHead
	for head != nil && head.Next != nil {
		node1 := head.Next
		node2 := head

		node2.Next = node1.Next
		node1.Next = node2

		pre.Next = node1
		pre = node2
		head = node2.Next
	}
	return dummyHead.Next
}
// @lc code=end

