/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 两区间无重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			if cur[1] > prev[1] {
				prev[1] = cur[1]
			}
		}
	}
	return append(res, prev)
}
// @lc code=end

