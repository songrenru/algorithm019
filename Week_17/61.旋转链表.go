/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	// 统计链表长度，并确定尾节点
	var l int
	tmp := head
	var tail *ListNode
	for tmp != nil {
		l++
		if tmp.Next == nil {
			tail = tmp
		}
		tmp = tmp.Next
	}

	k %= l
	if k == 0 {
		return head
	}

	// 确定断开点
	target := l - k
	tmp = head
	for i := 1; i < target; i++ {
		tmp = tmp.Next
	}

	// 重组链表
	newHead := tmp.Next
	tmp.Next = nil
	tail.Next = head

	return newHead
}

// @lc code=end

