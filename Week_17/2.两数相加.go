/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	up := 0
	dummyHead := &ListNode{}
	pre := dummyHead
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + up
		if sum >= 10 {
			up = 1
			sum -= 10
		} else {
			up = 0
		}
		pre.Next = &ListNode{Val: sum}
		pre = pre.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + up
		if sum >= 10 {
			up = 1
			sum -= 10
		} else {
			pre.Next = l1
			break
		}
		pre.Next = &ListNode{Val: sum}
		pre = pre.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + up
		if sum >= 10 {
			up = 1
			sum -= 10
		} else {
			pre.Next = l2
			break
		}
		pre.Next = &ListNode{Val: sum}
		pre = pre.Next
		l2 = l2.Next
	}
	if up == 1 {
		pre.Next = &ListNode{Val: 1}
	}

	return dummyHead.Next
}

// @lc code=end

