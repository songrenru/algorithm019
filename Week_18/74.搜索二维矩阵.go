/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// 竖向做二分查找
	// 很想做二分查找
	// 难点：竖向搜索时，取>=那一行
	if len(matrix) == 1 {
		return binarySearch(matrix[0], target)
	}
	row := binarySearchY(matrix, target)
	if matrix[row][0] > target && row > 0 {
		return binarySearch(matrix[row-1], target)
	}
	return binarySearch(matrix[row], target)
}

func binarySearchY(matrix [][]int, target int) int {
	lo, hi := 0, len(matrix)-1
	for lo < hi { // lo < hi，就已经排除了mid == hi的情况
		mid := int(uint(lo+hi) >> 1)
		if matrix[mid][0] == target {
			hi = mid
		} else if target < matrix[mid][0] {
			hi = mid - 1
		} else { // matrix[mid][0] < target
			lo = mid + 1
		}
	}
	return lo
}

func binarySearch(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] == target {
			return true
		} else if target < nums[mid] {
			hi = mid - 1
		} else { // nums[mid] < target
			lo = mid + 1
		}
	}
	return false
}

// @lc code=end

