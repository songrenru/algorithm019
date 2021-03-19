/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	first := head
	for head != nil {
		if head.Next != nil && head.Next.Val == head.Val {
			head.Next = head.Next.Next // 防止对象游离？
			continue
		}
		head = head.Next
	}
	return first
}

// @lc code=end

