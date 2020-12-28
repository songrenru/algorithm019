/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	if k == 0 || k > n {
		return nil
	}
	res := [][]int{}
	var dfs func(n, k, max int, path []int)
	dfs = func(n, k, max int, path []int) {
		// recursion terminator
		// 剪枝 // 这里即排除长度，也排除n = 5的情况，4-5+1=0
		if  k - len(path) > max - n + 1 {
			return
		}
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// if n == max + 1 { // 已被前两个条件过滤
		// 	return
		// }
		// process current logic
		// dirll down
		// 这里用for循环也比较自然，组合
		path = append(path, n)
		dfs(n + 1, k, max, path)

		path = path[:len(path) - 1]
		dfs(n + 1, k, max, path)
		// revert status
	}

	dfs(1, k, n, []int{})
	return res
}


// @lc code=end

