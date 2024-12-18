package main

import (
	"aoc2024/utils"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strings"
)

type Spot struct {
	spot string
	path []string
}

func main() {
	file, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(file), "\n")

	maxx := 0
	maxy := 0
	falling := make([][]int, len(lines))
	for i, r := range lines {
		falling[i] = make([]int, 2)
		fmt.Sscanf(r, "%d,%d", &falling[i][0], &falling[i][1])
		if falling[i][0] > maxx {
			maxx = falling[i][0]
		}
		if falling[i][1] > maxy {
			maxy = falling[i][1]
		}
	}
	// fmt.Println("FALLING:", len(falling), falling)

	grid := make([][]string, maxy+1)
	for y := range maxy + 1 {
		grid[y] = make([]string, maxx+1)
		for x := range maxx + 1 {
			grid[y][x] = "."
		}
	}
	// fmt.Println("ORIG GRID:", maxy, maxx)
	// for _, r := range grid {
	// 	fmt.Println(strings.Join(r, ""))
	// }

	findPath := func(part int) []string {
		start := []int{0, 0}
		end := []int{maxy, maxx}
		startKey := fmt.Sprintf(utils.GridYXKey, start[1], start[0])

		spotLookup := []Spot{}
		startSpot := Spot{
			spot: startKey,
			path: []string{startKey},
		}
		// Instead of puttin the Spot directly into the PriorityQueue, we'll add the index to a lookup table
		// ie. Id would be len(spotLookup)-1 immediately after appending to spotLookup.
		spotLookup = append(spotLookup, startSpot)

		gScores := make(map[string]int, len(grid)*len(grid[0])) // Part 2
		gScores[startKey] = 0
		// fScores := make(map[string]int, len(grid)*len(grid[0])) // Part 2
		h := func(spot Spot) int {
			var spotY, spotX int
			fmt.Sscanf(spot.spot, utils.GridYXKey, &spotY, &spotX)
			return int(math.Abs(float64(end[1]-spotY)) + math.Abs(float64(end[0]-spotX)))
		}
		// fScores[startKey] = h(startSpot)

		// goodPaths := utils.NewSet() // Part 2
		lowestScore := math.MaxInt
		shortestPath := []string{}

		// queue := []Spot{}
		// queue = append(queue, startSpot)
		// queuePos := 0 // Using manual pointer to iterate. PROS: Fast, efficient. CONS: No sort. :(
		queue := make(utils.PriorityQueue, 0)
		heap.Push(&queue, &utils.Item{Priority: 0, Value: 0})

		// for len(queue) > queuePos {
		for queue.Len() > 0 {
			// currQ := queue[queuePos]
			// queuePos += 1
			qItem := heap.Pop(&queue).(*utils.Item)
			currQ := spotLookup[qItem.Value]
			// fmt.Println("CURRQ:", qItem.Priority)

			var currY, currX int
			fmt.Sscanf(currQ.spot, utils.GridYXKey, &currY, &currX)
			// currName := grid[currY][currX]
			// fmt.Println("CURRNAME:", currName)

			if currY == end[0] && currX == end[1] {
				if part == 2 {
					return currQ.path
				}

				if len(currQ.path) < lowestScore {
					// fmt.Println("NEW LOW:", len(currQ.path))
					lowestScore = len(currQ.path)
					shortestPath = currQ.path
					continue
				}
			}

			for _, n := range utils.GridNeighbors {
				newY := currY + n[0]
				newX := currX + n[1]
				// fmt.Println("POS:", currY, currX, newY, newX, newName, nKey)

				if utils.IsValidPosInGrid(newY, newX, grid) {
					newName := grid[newY][newX]
					nKey := fmt.Sprintf(utils.GridYXKey, newY, newX)

					if newName != "#" && !slices.Contains(currQ.path, nKey) {
						tempG := gScores[currQ.spot] + 1
						newKeyG := fmt.Sprintf(utils.GridYXKey, newY, newX)
						newG, ok := gScores[newKeyG]
						if !ok {
							newG = math.MaxInt
						}
						if tempG < newG {
							if len(currQ.path) < lowestScore {
								if tempG < lowestScore {
									newPath := make([]string, len(currQ.path))
									copy(newPath, currQ.path)
									newPath = append(newPath, nKey)

									newSpot := Spot{
										spot: nKey,
										path: newPath,
									}
									spotLookup = append(spotLookup, newSpot)

									gScores[newKeyG] = tempG

									fScore := tempG + h(newSpot)
									// fScores[newKeyG] = fScore

									// fmt.Println("ADDING:", newSpot)
									// queue = append(queue, newSpot)
									heap.Push(&queue, &utils.Item{Priority: fScore, Value: len(spotLookup) - 1})
								}
							}
						}
					}
				}
			}
		}
		// return lowestScore - 1 // Disregard last square, we are counting steps
		if lowestScore == math.MaxInt {
			return nil
		}
		return shortestPath
	}

	// // New copy of grid
	// grid = make([][]string, len(origGrid))
	// for y := range grid {
	// 	grid[y] = append(grid[y], origGrid[y]...)
	// }

	// for _, f := range falling[:12] {
	for _, f := range falling[:1024] {
		grid[f[1]][f[0]] = "#"
	}
	// fmt.Println("GRID:", maxy, maxx)
	// for _, r := range grid {
	// 	fmt.Println(strings.Join(r, ""))
	// }
	res := findPath(1)
	fmt.Println("Part1:", len(res))

	for _, f := range falling[1025:] {
		grid[f[1]][f[0]] = "#" // Add new obstacle
		fKey := fmt.Sprintf(utils.GridYXKey, f[1], f[0])

		// Only rerun the search if the new wall interferes with the last path
		if slices.Contains(res, fKey) {
			res = findPath(2)
			if res == nil {
				rKey := fmt.Sprintf(utils.GridYXKey, f[0], f[1]) // XY format
				fmt.Println("Part2:", rKey)
				break
			}
			// fmt.Println("RES:", i+1025, res, f)
		}
	}
}
