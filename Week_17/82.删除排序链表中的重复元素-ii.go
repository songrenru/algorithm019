/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
	// 存在重复元素，返回下一个
	// 不存在，返回当前
	dummyHead := new(ListNode)
	pre := dummyHead
	for head != nil {
		dupl := false
		next := head.Next
		if next != nil && head.Val == next.Val {
			dupl = true
		}

		if dupl {
			for next != nil && next.Val == head.Val {
				next = next.Next
			}
			if next == nil {
				pre.Next = nil
				break
			}
			head = next
			// continue
		} else {
			pre.Next = head
			head = head.Next
			pre = pre.Next
		}
	}
	return dummyHead.Next
}

// @lc code=end

