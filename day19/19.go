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

	towels := strings.Split(sections[0], ", ")
	patterns := strings.Split(sections[1], "\n")
	// fmt.Println("T:", towels, "P:", patterns)

	// Tried sorting for Part 1, but after rewriting for Part 2, wasn't relevant.
	// byLengthThenAlphabetically := func(i int, j int) bool {
	// 	x := towels[i]
	// 	y := towels[j]
	// 	// deltaLength := len(x) - len(y)  // Short to Long
	// 	deltaLength := len(y) - len(x)  // Long to Short
	// 	return deltaLength < 0 || (deltaLength == 0 && x < y)
	// }
	// sort.Slice(towels, byLengthThenAlphabetically)
	// fmt.Println("Sorted T:", towels)

	var checkToken func(testStr string) int
	var checkTokenMem func(input string) int

	checkToken = func(testStr string) int {
		// fmt.Println("checkToken["+strconv.Itoa(d)+"]: testStr:", testStr)
		found := 0
		for it := 0; it < len(towels); it++ {
			t := towels[it]
			// fmt.Println("TESTING["+strconv.Itoa(d)+"]:", t, "TS:", testStr)
			if strings.HasPrefix(testStr, t) {
				newTestStr := testStr[len(t):]
				if len(newTestStr) > 0 {
					// fmt.Println("FOUND:", t, "ADJUSTED:", testStr)
					// newFound := checkToken(newTestStr)
					newFound := checkTokenMem(newTestStr)
					if newFound > 0 {
						found += newFound
						// break  // For Part 1
					}
				} else {
					found += 1
				}
			}
		}
		// fmt.Println("checkToken["+strconv.Itoa(d)+"]: found:", found)
		return found
	}
	checkTokenMem = utils.Memorized(checkToken)

	p1Total := 0
	p2Total := 0
	for _, p := range patterns {
		fmt.Println("PATTERN:", p)
		testStr := p[:]
		num := checkTokenMem(testStr)

		if num > 0 {
			p1Total += 1
			// fmt.Println("NUM:", num)
			p2Total += num
		}
	}

	fmt.Println("Part1:", p1Total)
	fmt.Println("Part2:", p2Total)
}
