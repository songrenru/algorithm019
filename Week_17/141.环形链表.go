/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// map,空间O(n)
	// 快慢指针，空间O(1)
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for {
		if slow == fast {
			return true
		}
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next

	}
}

// @lc code=end

