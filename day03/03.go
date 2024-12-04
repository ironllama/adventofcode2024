package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// input, err := os.ReadFile("03.sam2") // Sample is different! Tricky!
	input, err := os.ReadFile("03.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1 := func() int {
		// r := regexp.MustCompile(`(?<=mul\()\d+,\d+(?=\))`) // go regex doesn't support lookahead or lookbehind?!
		r := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matches := r.FindAllString(string(input), -1)
		// log.Println(matches)

		prod := 0
		for _, s := range matches {
			// justNums := strings.Split(s[4:len(s)-1], ",")
			// var allNums []int
			// for _, s := range(justNums) {
			// 	n, _ := strconv.Atoi(s)
			// 	allNums = append(allNums, n)
			// }
			// log.Println(justNums)

			// Same as above, but using fmt.Sscanf to use the format to assign values.
			var a, b int
			fmt.Sscanf(s[4:len(s)-1], "%d,%d", &a, &b)

			prod += a * b
		}
		return prod
	}

	part2 := func() int {
		r := regexp.MustCompile(`mul\(\d+,\d+\)|do[n't]*\(\)`)
		matches := r.FindAllString(string(input), -1)

		prod := 0
		doProd := true
		for _, s := range matches {
			if strings.HasPrefix(s, "do") {
				if strings.HasPrefix(s, "don't") {
					doProd = false
				} else {
					doProd = true
				}
				continue
			}

			if doProd {
				var a, b int
				fmt.Sscanf(s[4:len(s)-1], "%d,%d", &a, &b)

				prod += a * b
			}
		}

		return prod
	}

	log.Println("Part1:", part1())
	log.Println("Part2:", part2())
}
