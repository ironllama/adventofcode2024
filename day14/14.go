package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(file), "\n")

	// Read file and get inputs and fill setup variables
	maxx := 0
	maxy := 0
	robots := make([][]int, len(lines))
	for i, l := range lines {
		var px, py, vx, vy int
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		// fmt.Println("LINE:", px, py, vx, vy)

		if (px + 1) > maxx {
			maxx = px + 1
		}
		if (py + 1) > maxy {
			maxy = py + 1
		}

		robots[i] = []int{px, py, vx, vy}
	}
	// fmt.Println("ROBOTS:", robots)

	// Intialize the grid to keep track of robot positions
	grid := make([][]int, maxy)
	clearGrid := func() {
		for y := range maxy {
			grid[y] = make([]int, maxx)
		}
	}
	clearGrid()

	// _showGrid := func(title string) {
	// 	fmt.Println(title)
	// 	for _, r := range grid { // Print out grid
	// 		for _, c := range r {
	// 			fmt.Print(strconv.Itoa(c))
	// 		}
	// 		fmt.Println()
	// 	}
	// }
	// _showGrid("GRID START")

	// Figure out where each robot is after x seconds
	tick := func(secs int) {
		for _, r := range robots {
			newdx := r[2] * secs
			newdy := r[3] * secs

			newdx %= maxx // Compensate for overflow -- cycle back
			newdy %= maxy

			if newdx < 0 { // Compensate for underflow -- cycle back
				newdx = maxx + newdx
			}
			if newdy < 0 {
				newdy = maxy + newdy
			}

			newpx := r[0] + newdx
			newpy := r[1] + newdy

			newpx %= maxx
			newpy %= maxy

			grid[newpy][newpx] += 1
		}
	}

	// Part 1
	tick(100)
	// _showGrid("AFTER TICKS")

	countQuads := func(fromy int, toy int, fromx int, tox int) int {
		subtotal := 0
		for i := fromy; i < toy; i++ {
			for k := fromx; k < tox; k++ {
				subtotal += grid[i][k]
			}
		}
		return subtotal
	}
	halfx := maxx / 2
	halfy := maxy / 2

	total := 1
	total *= countQuads(0, halfy, 0, halfx)           // Top left
	total *= countQuads(0, halfy, halfx+1, maxx)      // Top right
	total *= countQuads(halfy+1, maxy, 0, halfx)      // Bottom left
	total *= countQuads(halfy+1, maxy, halfx+1, maxx) // Bottom right
	// numRobots := 0
	fmt.Println("Part1:", total)

	// Part 2
	num := 0
outer:
	for range 10000 {
		clearGrid()

		num += 1
		tick(num)

		for _, r := range grid {
			strLine := ""
			for _, c := range r {
				strLine += strconv.Itoa(c)
			}

			if strings.Contains(strLine, "11111111") {
				break outer
			}
		}
	}
	// _showGrid("TREE!")
	fmt.Println("Part 2:", num)
}
