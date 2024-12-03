package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read in the file...
	// file, err := os.Open("01.sam")
	file, err := os.Open("01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file) // Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line) // Fields splits by any whitespace

		if len(tokens) == 2 {
			num1, err1 := strconv.Atoi(tokens[0]) // Str to int
			num2, err2 := strconv.Atoi(tokens[1])
			if err1 != nil || err2 != nil {
				log.Fatal("Error converting numbers:", err1, err2)
				continue
			}

			left = append(left, num1)
			right = append(right, num2)
		}
	}

	sort.Ints(left[:])
	sort.Ints(right[:])

	// Find the diffs and add them up
	var sum int = 0
	for i := 0; i < len(left); i++ {
		d := int(math.Abs(float64((right[i] - left[i])))) // math.Abs operates on floats
		sum += d
	}

	log.Println("FINAL:", sum)
}
