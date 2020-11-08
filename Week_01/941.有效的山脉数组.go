/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

// @lc code=start
func validMountainArray(A []int) bool {
    i, lens := 0, len(A) - 1

    for ; i < lens && A[i] < A[i+1]; i++ {
    }
    if i == 0 || i == lens {
        return false
    }
    for ; i < lens && A[i] > A[i+1]; i++ {
    }
    return i == lens
}
// @lc code=end

