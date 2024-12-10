package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(file), "\n")

	ones := [][]int{}
	// Find all the 0's
	for y, l := range lines {
		for x, c := range l {
			if c == '0' { // runes!
				ones = append(ones, []int{y, x})
			}
		}
	}
	// log.Println("ONES:", len(ones), ones)

	type Spot struct {
		pos []int
		// visited []string   // Turns out to be unnecessary for this puzzle, since direction is always restricted to new value.
	}

	neighbors := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	part1 := func() {
		queue := []Spot{} // Since visited is unnecessary, the queue could have been simplified to just contain []int's
		total := 0
		for _, one := range ones {
			// queue = append(queue, Spot{one, []string{}})
			queue = append(queue, Spot{one})

			ninePath := []string{}
			for len(queue) > 0 {
				curr := queue[len(queue)-1]  // get from end of queue
				queue = queue[:len(queue)-1] // remove last elem from queue

				// Copy the visted and append the current pos
				// newVisited := make([]string, len(curr.visited))
				// copy(newVisited, curr.visited)
				// newVisited = append(newVisited, fmt.Sprintf("%d,%d", curr.pos[0], curr.pos[1]))

				currVal := int(lines[curr.pos[0]][curr.pos[1]] - '0') // Weird rune to int conversion (the inner rune - '0' produces an int32, so cast to general int)
				if currVal == 9 {
					thisNine := fmt.Sprintf("%d,%d", curr.pos[0], curr.pos[1])
					if slices.Index(ninePath, thisNine) == -1 {
						ninePath = append(ninePath, thisNine)
					}
				}
				nextVal := currVal + 1

				for _, n := range neighbors {
					newPos := []int{curr.pos[0] + n[0], curr.pos[1] + n[1]}

					if newPos[0] >= 0 && newPos[0] < len(lines) && newPos[1] >= 0 && newPos[1] < len(lines[0]) { // Inside valid area
						// if slices.Index(newVisited, fmt.Sprintf("%d,%d", newPos[0], newPos[1])) == -1 { // Avoid repeating visits (probably not necessary?)
						neighVal := int(lines[newPos[0]][newPos[1]] - '0')

						if neighVal == nextVal { // What we're looking for!
							// queue = append(queue, Spot{newPos, newVisited})
							queue = append(queue, Spot{newPos})
						}
						// }
					}
				}
			}
			total += len(ninePath)
		}
		log.Println("Part1:", total)
	}
	part1()

	part2 := func() {
		queue := [][]int{} // Since visited is unnecessary, the queue could have been simplified to just contain []int's
		total := 0
		for _, one := range ones {
			queue = append(queue, one)

			num := 0
			for len(queue) > 0 {
				curr := queue[len(queue)-1]  // get from end of queue
				queue = queue[:len(queue)-1] // remove last elem from queue

				currVal := int(lines[curr[0]][curr[1]] - '0') // Weird rune to int conversion (the inner rune - '0' produces an int32, so cast to general int)
				if currVal == 9 {
					// This is the only part that changes -- when finding the end, just count all of them.
					num += 1
				}
				nextVal := currVal + 1

				for _, n := range neighbors {
					newPos := []int{curr[0] + n[0], curr[1] + n[1]}

					if newPos[0] >= 0 && newPos[0] < len(lines) && newPos[1] >= 0 && newPos[1] < len(lines[0]) { // Inside valid area
						neighVal := int(lines[newPos[0]][newPos[1]] - '0')

						if neighVal == nextVal { // What we're looking for!
							queue = append(queue, newPos)
						}
					}
				}
			}
			total += num
		}
		log.Println("Part2:", total)
	}
	part2()
}
