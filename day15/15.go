package main

import (
	"aoc2024/utils"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := io.ReadAll(os.Stdin)
	sections := strings.Split(string(file), "\n\n")

	origGridStrs := strings.Split(sections[0], "\n")
	origGrid := make([][]string, len(origGridStrs))
	for y := range origGrid {
		origGrid[y] = strings.Split(origGridStrs[y], "")
	}

	inst := strings.ReplaceAll(sections[1], "\n", "")

	// Part 1
	grid := make([][]string, len(origGrid))
	for y := range grid {
		// grid[y] = make([]string, len(origGrid[0]))
		grid[y] = append(grid[y], origGrid[y]...)
	}

	ry := 0
	rx := 0
outer:
	for y, r := range grid {
		for x, c := range r {
			if string(c) == "@" {
				ry = y
				rx = x
				break outer
			}
		}
	}
	// fmt.Println("ROBOT:", ry, rx)

	var checkBox func(y int, x int, dy int, dx int) bool
	checkBox = func(y int, x int, dy int, dx int) bool {
		c := string(grid[y+dy][x+dx])
		if c == "." {
			grid[y][x], grid[y+dy][x+dx] = grid[y+dy][x+dx], grid[y][x]
			return true
		} else if c == "O" {
			good := checkBox(y+dy, x+dx, dy, dx)
			if good {
				grid[y][x], grid[y+dy][x+dx] = grid[y+dy][x+dx], grid[y][x]
				return true
			}
		}
		return false
	}

	for _, i := range inst {
		good := false

		dy := 0
		dx := 0
		if string(i) == "<" {
			dx = -1
		} else if string(i) == ">" {
			dx = 1
		} else if string(i) == "^" {
			dy = -1
		} else if string(i) == "v" {
			dy = 1
		}

		good = checkBox(ry, rx, dy, dx)
		if good {
			ry = ry + dy
			rx = rx + dx
		}
	}

	total1 := 0
	for y, r := range grid {
		for x, c := range r {
			if c == "O" {
				total1 += (y * 100) + x
			}
		}
	}
	fmt.Println("Part1:", total1)

	// Part 2
	newGrid := make([][]string, len(origGrid))
	for y := range newGrid {
		// newGrid[y] = make([]string, len(origGrid[0])*2)
		for x, c := range origGrid[y] {
			if c == "#" {
				newGrid[y] = append(newGrid[y], "#")
				newGrid[y] = append(newGrid[y], "#")
			} else if c == "." {
				newGrid[y] = append(newGrid[y], ".")
				newGrid[y] = append(newGrid[y], ".")
			} else if c == "O" {
				newGrid[y] = append(newGrid[y], "[")
				newGrid[y] = append(newGrid[y], "]")
			} else if c == "@" {
				newGrid[y] = append(newGrid[y], "@")
				newGrid[y] = append(newGrid[y], ".")
				ry = y
				rx = x * 2
			}
		}
	}
	// fmt.Println("ROBOT:", ry, rx, "NEW GRID:")
	// for _, r := range newGrid {
	// 	fmt.Println(strings.Join(r, ""))
	// }

	var checkBox2 func(y int, x int, dy int, dx int, n string) []string
	checkBox2 = func(y int, x int, dy int, dx int, n string) []string {
		if n == "." { // If it's clear, return a new set of swaps
			s := fmt.Sprintf("%d,%d,%d,%d", y, x, y+dy, x+dx)
			set := []string{s}
			return set
		} else if n == "[" || n == "]" {
			if dy == 0 { // Left and right just recurses the moves, like Part 1
				nN := newGrid[y+dy+dy][x+dx+dx]
				res := checkBox2(y+dy, x+dx, dy, dx, nN)
				if res != nil {
					s := fmt.Sprintf("%d,%d,%d,%d", y, x, y+dy, x+dx)
					res = append(res, s)
					return res
				}
			} else { // Oh boy...
				if n == "[" {
					// Test downstream's left and right for an up/down action
					nL := newGrid[y+dy+dy][x+dx+dx]
					goodL := checkBox2(y+dy, x+dx, dy, dx, nL)
					nR := newGrid[y+dy+dy][x+dx+dx+1]
					goodR := checkBox2(y+dy, x+dx+1, dy, dx, nR)

					if goodL != nil && goodR != nil { // If clear from both sides downstream
						res := append(goodL, goodR...)
						s := fmt.Sprintf("%d,%d,%d,%d", y, x, y+dy, x+dx)
						res = append(res, s)

						// Test to see if we need to swap upstream, based on matching box symbols
						if newGrid[y][x+1] == newGrid[y+dy][x+dx+1] { // If it's the same underneath
							s2 := fmt.Sprintf("%d,%d,%d,%d", y, x+1, y+dy, x+dx+1)
							res = append(res, s2)
						}

						return res
					}
				} else if n == "]" {
					// Test downstream's left and right for an up/down action
					nL := newGrid[y+dy+dy][x+dx+dx-1]
					goodL := checkBox2(y+dy, x+dx-1, dy, dx, nL)
					nR := newGrid[y+dy+dy][x+dx+dx]
					goodR := checkBox2(y+dy, x+dx, dy, dx, nR)

					if goodL != nil && goodR != nil { // If clear from both sides downstream
						res := append(goodL, goodR...)
						s := fmt.Sprintf("%d,%d,%d,%d", y, x, y+dy, x+dx)
						res = append(res, s)

						// Test to see if we need to swap upstream, based on matching box symbols
						if newGrid[y][x-1] == newGrid[y+dy][x+dx-1] { // If it's the same underneath
							s2 := fmt.Sprintf("%d,%d,%d,%d", y, x-1, y+dy, x+dx-1)
							res = append(res, s2)
						}

						return res
					}
				}
			}
		}
		return nil
	}

	// Go through each instruction
	for _, i := range inst {
		dy := 0
		dx := 0
		if string(i) == "<" {
			dx = -1
		} else if string(i) == ">" {
			dx = 1
		} else if string(i) == "^" {
			dy = -1
		} else if string(i) == "v" {
			dy = 1
		}

		c := newGrid[ry+dy][rx+dx]
		// fmt.Println("CHECK:", ry, rx, dy, dx, c, newGrid[ry][rx])
		res := checkBox2(ry, rx, dy, dx, c)
		if res != nil {
			res = utils.SliceRemoveDuplicatesFrom(res)
			// fmt.Println("MOVES:", res)
			for _, r := range res {
				// sargs := strings.Split(r, ",")
				var sya, sxa, syb, sxb int
				fmt.Sscanf(r, "%d,%d,%d,%d", &sya, &sxa, &syb, &sxb)

				// Prevent boxes next to the robot from moving. Weird edge cases?
				if string(i) == "^" || string(i) == "v" {
					if newGrid[ry][rx+1] == "[" && sya == ry && sxa == rx+1 {
						continue
					} else if newGrid[ry][rx-1] == "]" && sya == ry && sxa == rx-1 {
						continue
					}
				}

				newGrid[sya][sxa], newGrid[syb][sxb] = newGrid[syb][sxb], newGrid[sya][sxa]
			}

			// Move the robot!
			ry = ry + dy
			rx = rx + dx
		}

		// if x == 21 {
		// 	fmt.Println("STOPPED AT:", x, string(i))
		// 	break
		// }
	}

	total2 := 0
	for y, r := range newGrid {
		for x, c := range r {
			if c == "[" && newGrid[y][x+1] == "]" {
				total2 += (y * 100) + x
			}
		}
	}
	// fmt.Println("NEWGRID:")
	// for _, r := range newGrid {
	// 	fmt.Println(strings.Join(r, ""))
	// }
	fmt.Println("Part2:", total2)
}
