/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
	lens := len(nums)
	set := make(map[[4]int]bool)
	
    for i := 0; i < lens - 3; i++ {
		for j := i + 1; j < lens - 2; j++ {
			for k := j + 1; k < lens - 1; k++ {
				for l := k + 1; l < lens; l++ {
					temp := nums[i] + nums[j] + nums[k] + nums[l]
					if temp == target {
						set[[4]int{nums[i] ,nums[j] ,nums[k] ,nums[l]}] = true
					}
				}
			}
		}
	}
	res := [][]int{}
	for v, _ := range set {
		temp := append([]int{}, v[:]...)
		res = append(res, temp)
	}
	return res
}
// @lc code=end

