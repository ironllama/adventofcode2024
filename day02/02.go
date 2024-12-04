package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func arrayDelete(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func arrayCopy(orig []int) []int {
	newNums := make([]int, len(orig)) // Create a dynamically sized view
	copy(newNums, orig[:])            // Run the copy
	return newNums
}

func main() {
	// file, err := os.Open("02.sam")
	file, err := os.Open("02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Converted body of 02a
	arraySafe := func(arr []int) bool {
		// log.Println("ARR:", arr)
		sign := 0
		for i := 0; i < (len(arr) - 1); i++ {
			diff := arr[i] - arr[i+1]

			if sign == 0 {
				if diff > 0 {
					sign = 1
				} else {
					sign = -1
				}
			}

			if diff == 0 ||
				math.Abs(float64(diff)) > 3 ||
				(sign < 0 && diff > 0) ||
				(sign > 0 && diff < 0) {
				return false
			}
		}

		return true
	}

	var input [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		// log.Println("TOKENS:", tokens)

		// Convert string array to integer array
		var nums []int
		for _, s := range tokens {
			i, _ := strconv.Atoi(s) // Ignoring errs
			nums = append(nums, i)
		}

		input = append(input, nums)
	}

	part1 := func() int {
		numSafe := 0

		for _, nums := range input {
			if arraySafe(nums) {
				numSafe += 1
			}
		}

		return numSafe
	}

	part2 := func() int {
		numSafe := 0

	outer:
		for _, nums := range input {
			// Remove one from the array and test to see if it's safe
			// Only need one safe combo to continue to next line
			for i := 0; i < len(nums); i++ {
				newNums := arrayCopy(nums[:]) // Send as a view to send by ref?
				// log.Println("NEW:", newNums, "ORIG:", nums)
				if arraySafe(arrayDelete(newNums, i)) {
					// log.Println("SAFE")
					numSafe += 1
					continue outer
				}
			}
			// log.Println("UNSAFE")
		}

		return numSafe
	}

	log.Println("Part1:", part1())
	log.Println("Part2:", part2())
}
