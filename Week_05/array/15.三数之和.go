/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	lens := len(nums)
	for h := 0; h < lens - 2; h++ {
		if nums[h] > 0 {
			break
		}
		if h > 0 && nums[h] == nums[h-1] {
			continue
		}
		m, t := h + 1, lens - 1
		for m < t {
			sum := nums[h] + nums[m] + nums[t]
			if sum > 0 {
				for m < t {
					if nums[t] != nums[t-1] {
						t--
						break
					}
					t--
				}
			} else if sum < 0 {
				for m < t {
					if nums[m] != nums[m+1] {
						m++
						break
					}
					m++
				}
			} else {
				res = append(res, []int{nums[h], nums[m], nums[t]})
				for m < t {
					if nums[t] != nums[t-1] {
						t--
						break
					}
					t--
				}
				for m < t {
					if nums[m] != nums[m+1] {
						m++
						break
					}
					m++
				}
			}
		}
	}
	return res
}
// @lc code=end

