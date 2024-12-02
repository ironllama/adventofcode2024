package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func main() {
	// file, err := os.Open("02.sam")
	file, err := os.Open("02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numSafe := 0
	scanner := bufio.NewScanner(file)
	outerloop: for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		// log.Println("TOKENS:", tokens)

		// Convert string array to integer array
		var nums []int
		for _, s := range tokens {
			i, _ := strconv.Atoi(s)  // Ignoring errs, heh
			nums = append(nums, i)
		}

		sign := 0
		for i := 0; i < (len(nums) - 1); i++ {
			diff := nums[i] - nums[i+1]

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
					// log.Println("UNSAFE")
					continue outerloop
				}
		}
		// log.Println("SAFE")
		numSafe += 1
	}

	log.Println("FINAL:", numSafe)
}
