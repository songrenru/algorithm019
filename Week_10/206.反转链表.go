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
	// 利用额外空间的递归
	// 迭代
	// 递归
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	// newHead.Next = head // newHead只是作为返回结果，往上传递
	head.Next.Next = head // 传导，下一个节点指向当前节点
	head.Next = nil
	return newHead
}
// @lc code=end

