package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(input), "\n")

	grid := [][]string{}
	ants := make(map[string][][]int)
	for y, row := range lines {
		grid = append(grid, []string{})
		for x, col := range row {
			if col != '.' {
				_, ok := ants[string(col)]
				if !ok {
					ants[string(col)] = [][]int{}
				}
				ants[string(col)] = append(ants[string(col)], []int{y, x})
			}
			grid[y] = append(grid[y], string(col))
		}
	}
	// log.Println("ANTS:", ants, "GRID:", grid)

	part1 := func() {
		// Copy grid, in case of changes
		newGrid := [][]string{}
		for _, row := range grid {
			newRow := make([]string, len(row))
			copy(newRow, row)
			newGrid = append(newGrid, newRow)
		}
		// for _, row := range newGrid {  // Check the copy
		// 	log.Println(strings.Join(row, ""))
		// }

		count := 0
		for _, locs := range ants {
			for _, loc := range locs {
				for _, other := range locs {
					if loc[0] == other[0] && loc[1] == other[1] { // Ignore if this is the same point
						continue
					}
					diff := []int{loc[0] - other[0], loc[1] - other[1]}
					newLoc := []int{loc[0] + diff[0], loc[1] + diff[1]}
					// log.Println("A:", ant, "F:", loc, "T:", other, "D:", diff, "N:", newLoc, "G:")
					if newLoc[0] >= 0 && newLoc[0] < len(newGrid) && newLoc[1] >= 0 && newLoc[1] < len(newGrid[1]) {
						// if newGrid[newLoc[0]][newLoc[1]] == "." {  // Read the directions! Other antennas can be antinodes.
						if newGrid[newLoc[0]][newLoc[1]] != "#" {
							newGrid[newLoc[0]][newLoc[1]] = "#"
							count += 1
							// log.Println("A:", ant, "F:", loc, "T:", other, "D:", diff, "N:", newLoc)
						}
					}
				}
			}
		}
		// for _, row := range grid {  // Check visually against description.
		// 	log.Println(strings.Join(row, ""))
		// }
		log.Println("Part1:", count)
	}
	part1()

	part2 := func() {
		// for _, row := range grid { // Check to see if not reusing marked up map.
		// 	log.Println(strings.Join(row, ""))
		// }

		count := 0
		for _, locs := range ants {
			for _, loc := range locs {
				for _, other := range locs {
					if loc[0] == other[0] && loc[1] == other[1] { // Ignore if this is the same point
						continue
					}
					diff := []int{loc[0] - other[0], loc[1] - other[1]}
					newLoc := []int{loc[0] + diff[0], loc[1] + diff[1]}
					// log.Println("A:", ant, "F:", loc, "T:", other, "D:", diff, "N:", newLoc, "G:")
					for newLoc[0] >= 0 && newLoc[0] < len(grid) && newLoc[1] >= 0 && newLoc[1] < len(grid[1]) {
						if grid[newLoc[0]][newLoc[1]] != "#" {
							grid[newLoc[0]][newLoc[1]] = "#"
							// count += 1
							// log.Println("A:", ant, "F:", loc, "T:", other, "D:", diff, "N:", newLoc)
						}
						newLoc = []int{newLoc[0] + diff[0], newLoc[1] + diff[1]}
					}
				}
			}
		}
		for _, row := range grid {
			// log.Println(strings.Join(row, ""))  // Check visually against description
			for _, col := range row {
				// if col == "#" {  // Read carefully, dammit. Bitten twice!
				if col != "." {
					count += 1
				}
			}
		}
		log.Println("Part2:", count)
	}
	part2()
}
