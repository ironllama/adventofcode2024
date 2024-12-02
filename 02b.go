package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
)


func arrayDelete (slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
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

	scanner := bufio.NewScanner(file)
	numSafe := 0
	outer: for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		// log.Println("TOKENS:", tokens)

		// Convert string array to integer array
		var nums []int
		for _, s := range tokens {
			i, _ := strconv.Atoi(s)  // Ignoring errs
			nums = append(nums, i)
		}

		// Remove one from the array and test to see if it's safe
		// Only need one safe combo to continue to next line
		for i := 0; i < len(nums); i++ {
			newNums := make([]int, len(nums))  // Allocate destination array
			copy(newNums, nums[:])  // Run the copy
			// log.Println("NEW:", newNums, "ORIG:", nums)
			if (arraySafe(arrayDelete(newNums, i))) {
				// log.Println("SAFE")
				numSafe += 1
				continue outer
			}
		}
		// log.Println("UNSAFE")
	}

	log.Println("FINAL:", numSafe)
}