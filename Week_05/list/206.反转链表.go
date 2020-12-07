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
	// 递归
	if head == nil {
		return nil
	}
	
	dummyHead := new(ListNode)
	head = helper(head, dummyHead)
	head.Next = nil
	return dummyHead.Next
}

func helper(head, pre *ListNode) *ListNode {
	// terminator
	if head.Next == nil {
		pre.Next = head
		return head
	}
	// process current logic
	// drill down
	res := helper(head.Next, pre)
	res.Next = head

	return head
	// revert status if needed
}
// @lc code=end

