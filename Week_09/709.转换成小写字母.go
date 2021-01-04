/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 */

// @lc code=start
func toLowerCase(str string) string {
	arr := []byte(str)
    for i, c := range arr {
        if c >= 'A' && c <= 'Z' {
            arr[i] = c + 32
        }
    }
    return string(arr)
}
// @lc code=end

