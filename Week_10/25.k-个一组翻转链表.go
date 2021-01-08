/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 确定k group
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}

	// 先翻转下一个k group
	pre := reverseKGroup(tail, k)
	// 反转当前k区间
	for i := 0; i < k; i++ {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}
// @lc code=end

