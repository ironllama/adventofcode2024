package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// NOTE: This version was changed to read from STDIN.
// Will probably do this, going forward, instead of specifing a file.
func main() {
	// // file, _ := os.Open("05.sam")
	// file, _ := os.Open("05.txt")
	// defer file.Close()

	orders := make(map[string][]string)
	updates := [][]string{}

	readingOrders := true
	// scan := bufio.NewScanner(file)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()

		if readingOrders { // Read the first part: order pairings
			if line == "" {
				readingOrders = false
				continue
			}

			toks := strings.Split(line, "|")

			_, ok := orders[toks[0]]
			if !ok {
				orders[toks[0]] = []string{}
			}
			orders[toks[0]] = append(orders[toks[0]], toks[1])
		} else { // Read the second part: the updates to test
			toks := strings.Split(line, ",")

			updates = append(updates, toks)
		}
	}
	// log.Println("ORDERS:", orders)
	// log.Println("UPDATES:", updates)

	part1 := func() int {
		goodMiddles := 0

	outer:
		for i := range len(updates) {
			currUpdate := updates[i]
			// log.Println("CURR:", currUpdate)

			for k := range len(currUpdate) - 1 { // -1 - don't have to test last item
				for m := k + 1; m < len(currUpdate); m++ {
					valids, ok := orders[currUpdate[k]] // Test if map 'orders' contains
					if !ok || !slices.Contains(valids, currUpdate[m]) {
						// log.Println("KEY NOT FOUND:", currUpdate[k])
						continue outer
					}
				}
			}

			midIdx := int(math.Floor(float64(len(currUpdate)) / 2))
			midVal, _ := strconv.Atoi(currUpdate[midIdx])
			goodMiddles += midVal
		}

		return goodMiddles
	}

	part2 := func() int {
		goodMiddles := 0

		for i := range len(updates) {
			currUpdate := updates[i]

			bad := false
			for k := range len(currUpdate) - 1 {
				for m := k + 1; m < len(currUpdate); m++ {
					valids, ok := orders[currUpdate[k]] // Test if map 'orders' contains
					if !ok || !slices.Contains(valids, currUpdate[m]) {
						bad = true
						currUpdate[k], currUpdate[m] = currUpdate[m], currUpdate[k]
						m--
						continue
					}
				}
			}

			if bad {
				midIdx := int(math.Floor(float64(len(currUpdate)) / 2))
				midVal, _ := strconv.Atoi(currUpdate[midIdx])
				goodMiddles += midVal
			}
		}

		return goodMiddles
	}

	log.Println("Part1:", part1())
	log.Println("Part2:", part2())
}
