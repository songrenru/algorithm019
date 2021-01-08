/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	visited := map[*ListNode]bool{}
	for head != nil {
		if visited[head] {
			return head
		}
		visited[head] = true
		head = head.Next
	}
	return nil
}
// @lc code=end

