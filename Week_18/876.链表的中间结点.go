/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	low, fast := head, head
	for {
		if fast.Next == nil {
			return low
		}
		if fast.Next.Next == nil {
			return low.Next
		}
		low = low.Next
		fast = fast.Next.Next
	}
}

// @lc code=end

