/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	node := head
	for node != nil {
		count++
		node = node.Next
	}
	if count == n {
		return head.Next
	}
	skipNum := count - n
	pre := head
	for i := 1; i < skipNum; i++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return nil
	}
	pre.Next = pre.Next.Next
	return head
}
// @lc code=end

