/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 迭代
	// 递归
	if head == nil {
		return nil
	}
	dummyHead := new(ListNode)
	var helper func(node *ListNode) *ListNode
	helper = func(node *ListNode) *ListNode {
		if node.Next == nil {
			dummyHead.Next = node
			return node
		}
	
		newHead := helper(node.Next)
		newHead.Next = node
		return node
	}
	head = helper(head)
	head.Next = nil
	return dummyHead.Next
}


// @lc code=end

