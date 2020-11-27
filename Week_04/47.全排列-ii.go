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
	path := []int{}
	used := make([]bool, lens)

	var dfs func()
	dfs = func() {
		// recursion terminator
		if len(path) == lens {
			temp := make([]int, lens)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// process current logic
		// drill down
		pre := math.MinInt64
		for i := 0; i < lens; i++ {
			if !used[i] && pre != nums[i] {
				pre = nums[i]
				used[i] = true
				path = append(path, nums[i])

				dfs()

				used[i] = false
				path = path[:len(path)-1]

			}
		}
	}
	dfs()
	return res
}
// @lc code=end

