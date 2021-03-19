/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	// 基于递归的枚举
	lens := len(nums)
	res := [][]int{}

	var dfs func(i int, path []int)
	dfs = func(i int, path []int) {
		if i == lens {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		path = append(path, nums[i])
		dfs(i+1, path)

		path = path[:len(path)-1]
		dfs(i+1, path)
	}

	dfs(0, []int{})
	return res
	// 基于迭代的枚举
}

// @lc code=end

