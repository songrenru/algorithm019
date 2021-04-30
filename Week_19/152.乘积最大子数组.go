/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	// 暴力
	// dp
	// maxDp[i] = max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i], nums[i])
	// minDp[i] = min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i], nums[i])
	// 第三项是考虑前一项=0的情况
	l := len(nums)
	if l == 0 {
		return 0
	}

	mxF, mnF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < l; i++ {
		mx, mn := mxF, mnF
		mxF = max(mx*nums[i], max(mn*nums[i], nums[i]))
		mnF = min(mx*nums[i], min(mn*nums[i], nums[i]))
		ans = max(ans, mxF)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

