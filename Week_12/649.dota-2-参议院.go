/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	qR := []int{}
	qD := []int{}
	l := len(senate)
	for i := 0; i < l; i++ {
		if senate[i] == 'R' {
			qR = append(qR, i)
		} else {
			qD = append(qD, i)
		}
	}
	for len(qR) > 0 && len(qD) > 0 {
		if qR[0] < qD[0] {
			qR = append(qR, qR[0] + l)
		} else {
			qD = append(qD, qD[0] + l)
		}
		qR = qR[1:]
		qD = qD[1:]
	}
	if len(qR) > 0 {
		return "Radiant"
	}
	return "Dire"
}
// @lc code=end

