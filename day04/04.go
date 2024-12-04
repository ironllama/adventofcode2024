package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// file, _ := os.ReadFile("04.sam")
	file, _ := os.ReadFile("04.txt")
	lines := strings.Split(string(file), "\n")
	// log.Println("LINES:", lines)

	part1 := func() int {
		checkArm := func(y int, x int, dy int, dx int) bool {
			// log.Println("checkArm:", y, x, dy, dx)

			pattern := "MAS"
			for n := range 3 { // Range is 0-2
				newY := y + (dy * (n + 1))
				newX := x + (dx * (n + 1))
				// log.Println("CHECKING:", newY, newX)

				if newY < 0 || newY >= len(lines) || newX < 0 || newX >= len(lines[0]) {
					return false
				}

				// log.Println("TEST:", string(lines[newY][newX]), string(pattern[n]))
				if lines[newY][newX] != pattern[n] {
					return false
				}
			}

			// log.Println("FOUND:", y, x)
			return true
		}

		good := 0
		for y, line := range lines {
			for x, char := range line {
				// log.Println(y, x, string(char))
				if string(char) == "X" {
					for _, dy := range [3]int{-1, 0, 1} {
						for _, dx := range [3]int{-1, 0, 1} {
							if checkArm(y, x, dy, dx) {
								good += 1
							}
						}
					}
				}
			}
		}

		return good
	}

	part2 := func() int {
		checkDiags := func(y int, x int) bool {
			// log.Println("checkDiags:", y, x)
			// Make sure the middle is within 1 space of edges, since the X needs one space to each direction.
			if y < 1 || y >= (len(lines)-1) || x < 1 || x >= (len(lines[0])-1) {
				// log.Println("Out of bounds.")
				return false
			}

			uldr := string(lines[y-1][x-1]) + string(lines[y+1][x+1]) // up-left to down-right
			// log.Println("TESTING:", uldr)
			if !(uldr == "MS" || uldr == "SM") {
				// log.Println("NOPE")
				return false
			}

			urdl := string(lines[y-1][x+1]) + string(lines[y+1][x-1]) // up-right to down-left
			// log.Println("TESTING:", uldr)
			if !(urdl == "MS" || urdl == "SM") {
				// log.Println("NOPE")
				return false
			}

			return true
		}

		good := 0
		for y, line := range lines {
			for x, char := range line {
				// log.Println(y, x, string(char))
				if string(char) == "A" {
					if checkDiags(y, x) {
						good += 1
					}
				}
			}
		}

		return good
	}

	log.Println("Part1:", part1())
	log.Println("Part2:", part2())
}
