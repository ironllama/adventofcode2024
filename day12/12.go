package main

import (
	"aoc2024/utils"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(input), "\n")

	fkey := "%d,%d"
	neighbors := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} // Left first.
	allAround := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}, {1, -1}, {-1, -1}, {-1, 1}, {1, 1}}
	validPos := func(y int, x int) bool {
		return y >= 0 && y < len(lines) && x >= 0 && x < len(lines[0])
	}

	debug := false

	solve := func() {
		p1Total := 0
		p2Total := 0
		// enclosed := 0 // For debugging. Number of plots totally enclosed by another.

		// A reference structure to make sure we visit plots only once.
		toVisit := make([][]int, len(lines))
		for i := range toVisit {
			toVisit[i] = make([]int, len(lines[0]))
		}
		// fmt.Println("VISIT:", toVisit)

		// For Part 2
		allPlots := map[string][]string{}
		allSides := map[string][]int{}

		for y, row := range lines {
			for x := range row {
				if debug {
					fmt.Println("CHECKING:", y, x, string(lines[y][x]), toVisit[y][x], toVisit)
				}

				// Mostly Part 1
				if toVisit[y][x] == 1 {
					continue
				}

				area := 0
				borders := 0
				outerLetter := string(lines[y][x])

				// Using the queue as a "history" - need advancing cursor, rather than pop
				queuePos := 0
				queue := []string{}
				queue = append(queue, fmt.Sprintf(fkey, y, x))
				for queuePos < len(queue) {
					// fmt.Println("QUEUE:", queuePos, queue[queuePos], queue)
					currQueue := queue[queuePos]
					queuePos += 1

					var currY, currX int
					fmt.Sscanf(currQueue, fkey, &currY, &currX)

					pKey := outerLetter + fmt.Sprintf(fkey, y, x)
					allPlots[pKey] = append(allPlots[pKey], currQueue)

					toVisit[currY][currX] = 1
					area += 1

					for _, n := range neighbors {
						newY := currY + n[0]
						newX := currX + n[1]
						// fmt.Println("POS:", currY, currX, newY, newX)

						if validPos(newY, newX) {
							newLetter := string(lines[newY][newX])

							if newLetter == outerLetter {
								newKey := fmt.Sprintf(fkey, newY, newX)

								if !slices.Contains(queue, newKey) {
									queue = append(queue, newKey)
								}
							} else {
								borders += 1
							}
						} else {
							borders += 1
						}
					}
				}
				// fmt.Println(outerLetter, ":", area, "*", borders, "=", area*borders)
				p1Total += area * borders

				// Part 2 from here.
				sides := 1 // Advance at every turn (corner)
				dir := 0
				// neighborNames := []string{}
				neighborNames := utils.NewSet()

				if debug {
					fmt.Println("REUSING QUEUE:", queuePos, queue)
				}

				// Find furthest lower right corner to start.
				newPlotY := 0
				newPlotX := 0
				for _, q := range queue {
					var thisY, thisX int
					fmt.Sscanf(q, fkey, &thisY, &thisX)
					if thisY > newPlotY {
						newPlotY = thisY
						newPlotX = thisX
					} else if thisY == newPlotY && thisX > newPlotX {
						newPlotX = thisX
					}
				}

				// Bootstrap the first plot using the last plot processed above.
				firstBorderY := newPlotY
				firstBorderX := newPlotX + 1 // One right to bottom right has to be part of the border.

				if validPos(firstBorderX, firstBorderY) {
					// neighborNames = append(neighborNames, string(lines[firstBorderY][firstBorderX]))
					neighborNames.Add(string(lines[firstBorderY][firstBorderX]))
				} else {
					// neighborNames = append(neighborNames, fmt.Sprintf(fkey, firstBorderY, firstBorderX))
					neighborNames.Add(fmt.Sprintf(fkey, firstBorderY, firstBorderX))
				}

				// For edge case where the plot is shaped like a backwards L.
				// Since we're going from right to left, it will double back after finishing the base
				// but when encountering the start point, it is not yet finished.
				upFromStart := false
				if newPlotY > 0 && string(lines[newPlotY-1][newPlotX]) == outerLetter {
					upFromStart = true
				}

				if debug {
					fmt.Println("STARTING:", outerLetter, newPlotY, newPlotX, dir, neighbors[dir])
				}

				for {
					leftSide := neighbors[(dir+3)%4]
					leftPos := []int{newPlotY + leftSide[0], newPlotX + leftSide[1]}

					wallOnLeft := false

					// Check neighbors in all directions
					// Include diagonal for the edge case of being *almost* fully enclosed, except for diagonal.
					for _, d := range allAround {
						newY := newPlotY + d[0]
						newX := newPlotX + d[1]

						if !validPos(newX, newY) {
							// neighborNames = append(neighborNames, fmt.Sprintf(fkey, newY, newX))
							neighborNames.Add(fmt.Sprintf(fkey, newY, newX))
						} else if string(lines[newY][newX]) != outerLetter {
							// neighborNames = append(neighborNames, string(lines[newY][newX]))
							neighborNames.Add(string(lines[newY][newX]))
						}
					}

					if !validPos(leftPos[0], leftPos[1]) {
						if debug {
							fmt.Println("WALL FOUND OUT OF BOUNDS.")
						}
						wallOnLeft = true
						// neighborNames = append(neighborNames, fmt.Sprintf(fkey, leftPos[0], leftPos[1]))
						neighborNames.Add(fmt.Sprintf(fkey, leftPos[0], leftPos[1]))
					} else if string(lines[leftPos[0]][leftPos[1]]) != outerLetter {
						if debug {
							fmt.Println("WALL ON LEFT. STRAIGHT.")
						}
						wallOnLeft = true
						// neighborNames = append(neighborNames, string(lines[leftPos[0]][leftPos[1]]))
						neighborNames.Add(string(lines[leftPos[0]][leftPos[1]]))
					} else {
						if debug {
							fmt.Println("NO WALL. TURNING.")
						}
					}

					if debug {
						fmt.Println("LEFT:", leftSide, leftPos, wallOnLeft, neighborNames)
					}

					if !wallOnLeft {
						if debug {
							fmt.Println("NOTHING ON LEFT, STARTING RANGE WITH LEFT SIDE.")
						}
						dir = (dir + 3) % 4 // Start with the left, if there's nothing there.
						sides += 1          // Turn!
					}

					// If not border, turn until you can proceed. Adding 1 side for every turn.
					for i := range 4 {
						if i > 0 {
							sides += 1
							if debug {
								fmt.Println("TURNED: sides:", sides)
							}
						}

						newDir := (dir + i) % 4
						newY := newPlotY + neighbors[newDir][0]
						newX := newPlotX + neighbors[newDir][1]

						if validPos(newY, newX) {
							nextName := lines[newY][newX]
							if debug {
								fmt.Println("PROPOSING:", string(nextName), outerLetter, newY, newX)
							}

							if string(nextName) == outerLetter {
								dir = newDir
								newPlotY = newY
								newPlotX = newX
								break
							}
						}
					}

					if newPlotY == firstBorderY && newPlotX == (firstBorderX-1) { // -1 on X to negate the +1 added earlier
						// Back at start.
						if upFromStart && dir == 2 { // If coming from left, might be backwards L. Keep going left!
							upFromStart = !upFromStart
							continue
						} else if dir != 0 { // Rotate to face back as started to capture final wall.
							sides += 3 - dir
						}
						break
					}

					if debug {
						fmt.Println("MOVING FORWARD TO:", newPlotY, newPlotX, "STARTED FROM:", firstBorderY, firstBorderX-1)
					}
				}

				if sides%2 != 0 && dir == 0 {
					sides -= 1 // Sometimes counting does not start at lower right corder, so last side gets counted twice -- at beginning and end.
				}

				sKey := outerLetter + fmt.Sprintf(fkey, y, x) // Must align with pKey and allPlots!
				allSides[sKey] = append(allSides[sKey], []int{area, sides}...)

				if neighborNames.Size() == 1 {
					for k, v := range allPlots {
						if slices.Index(v, fmt.Sprintf(fkey, firstBorderY, firstBorderX)) != -1 {
							allSides[k][1] += sides
							// enclosed += 1
						}
					}
				}
			}
		}

		fmt.Println("Part1:", p1Total)

		// fmt.Println("PLOTS:", len(allPlots))
		for k, v := range allSides {
			if debug {
				fmt.Println("SIDES:", k, v, v[0]*v[1])
			}
			p2Total += v[0] * v[1]
		}
		fmt.Println("Part2:", p2Total)
	}
	solve()
}
