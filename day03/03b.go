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

	log.Println("FINAL:", prod)
}
