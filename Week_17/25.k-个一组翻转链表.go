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
	// 递归 + 组内翻转

	// 确定下一组的第一个元素
	kNode := head
	for i := 0; i < k; i++ {
		if kNode == nil {
			return head
		}
		kNode = kNode.Next
	}

	// 反转当前组的元素
	pre := reverseKGroup(kNode, k)
	for i := 0; i < k; i++ {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	// 返回新的链表头
	return pre
}

// @lc code=end

