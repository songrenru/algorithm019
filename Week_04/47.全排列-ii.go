/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	lens := len(nums)
	res := [][]int{}
	visited := make([]bool, lens)

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == lens {
			temp := make([]int, lens)
			copy(temp, path)
			res = append(res, temp)
			return
		}
	
		// pre := math.MinInt64
		for i := 0; i < lens; i++ {
			// if !used[i] && pre != nums[i] {
			// 	// 没用过，且不等于前一个或几个，才下探
			// 	// 可以做到平移排重，下探不排重
			// 	// 但是有跳跃性，不直观，不好理解
			// 	pre = nums[i]
			// 	...

			// }
			if i > 0 && !visited[i-1] && nums[i] == nums[i - 1] { 
				// 平移排重，下探不能排重
				/* !visited[i-1]是关键点
				 * 下探时，上一个相同值已visit
				 * 平移时，上一个相同值已经回溯恢复到了not visit状态
				 */
				continue
			}
			if visited[i] {
				continue
			}
	
			path = append(path, nums[i])
			visited[i] = true
			dfs(path)
	
			path = path[:len(path) - 1]
			visited[i] = false
			// dfs(path)
		}
	}
	dfs([]int{})
	return res
}
// @lc code=end

