/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	lens := len(nums)
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == lens {
			temp := make([]int, lens)
			copy(temp, path)
			res = append(res, temp)
			return
		}
	
		for i := 0; i < lens; i++ {
			if visited[nums[i]] {
				continue
			}
	
			path = append(path, nums[i])
			visited[nums[i]] = true
			dfs(path)
	
			path = path[:len(path) - 1]
			visited[nums[i]] = false
			// dfs(path)
		}
	}
	dfs([]int{})
	return res
}


// @lc code=end

