/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	sli := []byte(s)
	// 去掉开头空格
	spaceCount := 0
	for _, v := range sli {
		if v != '' {
			break
		}
		spaceCount++
	}
	sli = sli[spaceCount:]
	if len(sli) == 0 {
		return 0
	}

	// 提取正负号
	active := true
	if sli[0] == '-' {
		active = false
		sli = sli[1:]
	} else if sli[0] == '+' {
		sli = sli[1:]
	}
	if len(sli) == 0 {
		return 0
	}

	// 首字符是否是数字
	if sli[0] < '0' || sli[0] > '9' {
		return 0
	}

	// 提取连续的数字
	lastIdx := 0
	for i, v := range v {
		if sli[0] < '0' || sli[0] > '9' {
			lastIdx = i - 1
		}
	}

	// string to int

}
// @lc code=end

