package main

import (
	"aoc2024/utils"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strings"
)

type Spot struct {
	spot  string
	path  []string
	score int
	dir   int
	// dist  int
}

func main() {
	file, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(file), "\n")

	startX := 0
	startY := 0
	endX := 0
	endY := 0
	grid := [][]string{}
	for y, r := range lines {
		newRow := []string{}
		for x, c := range r {
			if string(c) == "S" {
				startY = y
				startX = x
			} else if string(c) == "E" {
				endY = y
				endX = x
			}
			newRow = append(newRow, string(c))
		}
		grid = append(grid, newRow)
	}
	// fmt.Println("GRID:", startY, startX, endY, endX)
	// for _, r := range grid {
	// 	fmt.Println(strings.Join(r, ""))
	// }

	lowestScore := math.MaxInt

	gScores := make(map[string]int, len(grid)*len(grid[0])) // Part 2
	// fScores := make(map[string]int, len(grid)*len(grid[0])) // Part 2
	// h := func(spot Spot) int {
	// 	var spotY, spotX int
	// 	fmt.Sscanf(spot.spot, utils.GridYXKey, &spotY, &spotX)
	// 	return int(math.Abs(float64(endY - spotY)) + math.Abs(float64(endX - spotX)))
	// }

	goodPaths := utils.NewSet() // Part 2

	queue := []Spot{}
	s := fmt.Sprintf(utils.GridYXKey, startY, startX)
	queue = append(queue, Spot{
		spot:  s,
		path:  []string{s},
		score: 0,
		dir:   1, // MAKE SURE YOU ALIGN THIS CORRECTLY w/ INSTRUCTIONS AND utils.GridNeighbors!!!
		// dist: int(math.Abs(float64(endY - startY)) + math.Abs(float64(endX - startX))),
	})
	queuePos := 0 // Using manual pointer to iterate. PROS: Fast, efficient. CONS: No sort. :(

	for len(queue) > queuePos {
		currQ := queue[queuePos]
		queuePos += 1
		// fmt.Println("CURRQ:", currQ)

		var currY, currX int
		fmt.Sscanf(currQ.spot, utils.GridYXKey, &currY, &currX)
		// currName := grid[currY][currX]
		// fmt.Println("CURRNAME:", currName)

		if currY == endY && currX == endX {
			// fmt.Println("END:", currQ)
			if currQ.score < lowestScore {
				lowestScore = currQ.score
				goodPaths = utils.NewSet()
			}

			if currQ.score == lowestScore {
				for _, p := range currQ.path {
					goodPaths.Add(p)
				}
			}
		}

		for ni, n := range utils.GridNeighbors {
			newY := currY + n[0]
			newX := currX + n[1]
			newName := grid[newY][newX]
			nKey := fmt.Sprintf(utils.GridYXKey, newY, newX)
			// fmt.Println("POS:", currY, currX, newY, newX, newName, nKey)

			if newName != "#" && !slices.Contains(currQ.path, nKey) {
				scoreAdjust := 1
				if ni != currQ.dir {
					scoreAdjust += 1000
				}
				tempG := currQ.score + scoreAdjust
				newKeyG := fmt.Sprintf(utils.GridYXKey+",ni", newY, newX, ni)
				newG, ok := gScores[newKeyG]
				if !ok {
					newG = math.MaxInt
				}
				if tempG <= newG {
					gScores[newKeyG] = tempG

					if tempG <= lowestScore {
						newPath := make([]string, len(currQ.path))
						copy(newPath, currQ.path)
						newPath = append(newPath, nKey)

						newSpot := Spot{
							spot:  nKey,
							path:  newPath,
							score: tempG,
							dir:   ni,
						}
						// fmt.Println("ADDING:", newSpot)
						queue = append(queue, newSpot)

						// newF := tempG + h(newSpot)
						// newF := scoreAdjust
						// newF := newSpot.score
						// fScore[nKey] = newF
						// newSpot.fScore = newF

						// var spotQueue Spot
						// for _, q := range queue {
						// 	if q.spot == nKey {
						// 		spotQueue = q
						// 		break
						// 	}
						// }
						// if spotQueue.spot != "" {  // No nils for structs, check default values on properties
						// 	spotQueue.path = newSpot.path
						// 	spotQueue.score = newSpot.score
						// 	spotQueue.dir = newSpot.dir
						// 	spotQueue.dist = newSpot.dist
						// 	spotQueue.fScore = newSpot.fScore
						// }
					}
				}
			}
		}
	}

	fmt.Println("Part1:", lowestScore)
	fmt.Println("Part2:", goodPaths.Size())

	// fmt.Println("FINAL:")
	// // for _, p := range goodPaths.List() {
	// // 	fmt.Println(p)
	// // }
	// for y, r := range grid {
	// 	for x, c := range r {
	// 		nKey := fmt.Sprintf(utils.GridYXKey, y, x)
	// 		if goodPaths.Contains(nKey) {
	// 			fmt.Print("X")
	// 		} else {
	// 			fmt.Print(c)
	// 		}
	// 	}
	// 	fmt.Println("")
	// }
}
