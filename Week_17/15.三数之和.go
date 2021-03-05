/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// 暴力 + 去重，O(n^3)
	// 排序 + 二分查找 + 去重，O(N^2 * logN)
	// 排序 + 双指针剪枝去重，O(N^2)
	sort.Ints(nums)
	l := len(nums)
	lastIdx1 := l - 2
	res := [][]int{}
	for i := 0; i < lastIdx1; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cmp := -nums[i]
		left, right := i+1, l-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == cmp {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > cmp {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			}
		}
	}
	return res
}

// @lc code=end

