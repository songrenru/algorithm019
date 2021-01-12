/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	l := len(candidates)
	res := [][]int{}
	var dfs func(idx, sum int, path []int, choice bool)
	dfs = func(idx, sum int, path []int, choice bool) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if idx == l {
			return
		}
		// 上一层不选，且该层(选择时)元素和上一层重复时，子树相同，剔除
		// 若该层也不选，则继续往下走
		if idx > 0 && candidates[idx] == candidates[idx - 1] && !choice {
			dfs(idx + 1, sum, path, false)
			return
		}
	
		sum += candidates[idx]
		path = append(path, candidates[idx])
		dfs(idx + 1, sum, path, true)

		sum -= candidates[idx]
		path = path[:len(path) - 1]
		dfs(idx + 1, sum, path, false)
	}
	dfs(0, 0, []int{}, true)
	return res
}


// @lc code=end

