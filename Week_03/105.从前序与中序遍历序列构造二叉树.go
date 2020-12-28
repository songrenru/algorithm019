/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// recursion terminator
	if len(preorder) == 0 {
		return nil
	}
	// process current logic
	// drill down
	root := &TreeNode{Val:preorder[0]}
	rootIdx := 0
	for i, v := range inorder {
		if v == preorder[0] {
			rootIdx = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:rootIdx])+1], inorder[:rootIdx])
	root.Right = buildTree(preorder[len(inorder[:rootIdx])+1:], inorder[rootIdx + 1:])
	return root
	// revert status if needed
}
// @lc code=end

