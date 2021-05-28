/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	r := []int{}
	d := []int{}
	l := len(senate)
	for i := 0; i < l; i++ {
		if senate[i] == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+l)
		} else { // >
			d = append(d, d[0]+l)
		}
		r = r[1:]
		d = d[1:]
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}

// @lc code=end

