package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	file, _ := io.ReadAll(os.Stdin)
	machines := strings.Split(string(file), "\n\n")

	p1Total := 0
	p2Total := 0

	for _, m := range machines {
		// fmt.Println("START M:", m)

		lines := strings.Split(m, "\n")
		var butAx, butAy, butBx, butBy, px, py int
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &butAx, &butAy)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &butBx, &butBy)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &px, &py)
		// fmt.Println("A:", butAx, butAy, "B:", butBx, butAy, "P:", px, py)

		// Part 1
		lowest := math.MaxInt32
		for na := range 100 {
			for nb := range 100 {
				nx := (na * butAx) + (nb * butBx)
				ny := (na * butAy) + (nb * butBy)

				if nx > px || ny > py {
					break
				}

				if nx == px && ny == py {
					cost := (na * 3) + nb
					if cost < lowest {
						lowest = cost
					}
					// fmt.Println("FOUND:", cost, lowest)
				}
			}
		}
		if lowest != math.MaxInt32 {
			p1Total += lowest
		}

		// Part 2
		px += 10000000000000
		py += 10000000000000
		// Use the point-slope formula using final point and one slope to get the y-interscept
		// and line equation. Then find the intersection with the other line going through 0,0
		// The point where they meet is used against the second line to find the number of times
		// for B and the left over is the number of times for A.
		// point-slope: y - y1 = m(x - x1)
		// if x = 0: y = (m * (0 - x1)) + y1
		slopeA := float64(butAy) / float64(butAx)
		s1y := (slopeA * (0 - float64(px))) + float64(py) // slop 1 y axis intercept
		// fmt.Println("s1y:", s1y)

		// Intersection of two lines:
		// x = (y2 - y1) / (m1 - m2)  // y1 is 0, since the B slope is measured from 0,0
		// y = (m1 * x) + b1
		// In this case, B uses m1 and y1
		slopeB := float64(butBy) / float64(butBx)
		ix := (s1y - 0) / (slopeB - slopeA) // Intersection x
		// iy := (slopeB * ix) + 0             // Intersection y using line equation for B
		// fmt.Println("ix:", ix)

		// Use intersection point to find how many times B and A were pressed
		numB := ix / float64(butBx)
		numA := (float64(px) - ix) / float64(butAx)

		numBint := math.Round(numB)
		numAint := math.Round(numA)
		diffB := math.Abs(numB - numBint)
		diffA := math.Abs(numA - numAint)
		// fmt.Println("NUMS: B:", numB, diffB, "A:", numA, diffA)
		// Some massaging, because of crappy computer floating-point math
		// .01 seems legit?!
		if diffB < .01 && diffA < .01 {
			// fmt.Println("GOOD VAL:", numA, numB, numAint, numBint)

			subtotal := (numAint * 3) + numBint
			p2Total += int(subtotal)
			// fmt.Println("SUBTOTAL:", subtotal, p2Total)
		}
	}

	fmt.Println("Part1:", p1Total)
	fmt.Println("Part2:", p2Total)
}
