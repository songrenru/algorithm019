/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre := dummyHead
	for pre.Next != nil {
		node := pre.Next
		if node.Val == val {
			next := node.Next
			pre.Next = next
			node.Next = nil
		} else {
			pre = node
		}
	}
	return dummyHead.Next
}
// @lc code=end

