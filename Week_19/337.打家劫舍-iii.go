/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var mem map[*TreeNode]int = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if money, ok := mem[root]; ok {
		return money
	}

	money := root.Val
	if root.Left != nil {
		money += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		money += rob(root.Right.Left) + rob(root.Right.Right)
	}
	mem[root] = max(money, rob(root.Left)+rob(root.Right))
	return mem[root]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

