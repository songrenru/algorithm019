/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	res := [][]int{}

	var dfs func(i int, path []int)
	dfs = func(i int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for j := i; j < l; j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			path = append(path, nums[j])
			dfs(j+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{})
	return res
}

// @lc code=end

