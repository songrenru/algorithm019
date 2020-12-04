/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	const (
		noth = iota
		east
		south
		west
	)
	obstaclesSet := make(map[[2]int]bool)
	for _, obstacle := range obstacles {
		obstaclesSet[[2]int{obstacle[0], obstacle[1]}] = true
	}

	res := 0
	direct := noth
	x, y := 0, 0
	for _, cmd := range commands {
		if cmd == -1 {
			direct = (direct + 1) % 4
			continue
		} else if cmd == -2 {
			direct = (direct + 3) % 4
			continue
		} else {
			for i := 0; i < cmd; i++ {
				yChg , xChg := 0, 0
				switch direct {
				case noth:
					y++
					yChg--
				case east:
					x++
					xChg--
				case south:
					y--
					yChg++
				case west:
					x--
					xChg++
				}
				if obstaclesSet[[2]int{x, y}] {
					x += xChg
					y += yChg
					break
				}
			}
		}
		distance := x*x + y*y
		if res < distance {
			res = distance
		}
	}

	return res
}
// @lc code=end

