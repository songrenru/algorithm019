/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
    max := 0
    i, j := 0, len(height)-1
    for i < j {
    	l, r := height[i], height[j]
		if cap := min(l, r) * (j - i); cap > max {
			max = cap
		}
		if l < r {
			i++
		} else {
			j--
		}
    }
    return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// m1:暴力破解
// m2:双指针
// @lc code=end

