/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 暴力枚举
	// dp，空间优化 // 空间优化后，和贪心法的代码形似差不多
	// dp[i] = max(dp[i-1], 0) + nums[i]
	/* 严格来说，这不是dp，因为还需要额外求max，dp[i]并不能代表截止到当前的最大子序和
	而只能代表“连续”到当前的最大子序和 */
	// 上面条注释，理解有误，dp是比较灵活的，不要太死脑筋

	// 分治法
	return get(nums, 0, len(nums)-1).mSum
}

func get(nums []int, l, r int) status {
	if l == r {
		return status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := int(uint(l+r) >> 1)
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func pushUp(l, r status) status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return status{lSum, rSum, mSum, iSum}
}

type status struct {
	lSum, rSum, mSum, iSum int
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

