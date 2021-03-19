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
	// map,空间O(n)
	// 快慢指针，数学推导
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head.Next, head.Next.Next
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		if slow == fast {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	for {
		if head == slow {
			return head
		}
		head = head.Next
		slow = slow.Next
	}
}

// @lc code=end

