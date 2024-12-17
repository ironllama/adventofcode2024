package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := io.ReadAll(os.Stdin)
	sections := strings.Split(string(file), "\n\n")
	registers := strings.Split(sections[0], "\n")
	programStr := strings.Split(sections[1], " ")
	programs := strings.Split(programStr[1], ",")

	var ra, rb, rc int
	fmt.Sscanf(registers[0], "Register A: %d", &ra)
	fmt.Sscanf(registers[1], "Register B: %d", &rb)
	fmt.Sscanf(registers[2], "Register C: %d", &rc)
	// fmt.Println("INPUT:", ra, rb, rc, programs)

	combo := func(oper int) int {
		if oper > 0 && oper <= 3 {
			return oper
		} else if oper == 4 {
			return ra
		} else if oper == 5 {
			return rb
		} else if oper == 6 {
			return rc
		} else if oper == 7 {
			fmt.Println("ILLEGAL:", oper)
		}
		fmt.Println("UNKNOWN:", oper)
		return -1
	}

	runProgram := func() string {
		iptr := 0
		output := ""
		for len(programs) > iptr {
			opCode, _ := strconv.Atoi(programs[iptr])
			iptr += 1
			operand, _ := strconv.Atoi(programs[iptr])
			iptr += 1

			if opCode == 0 { // adv
				ra = ra / int(math.Pow(2, float64(combo(operand))))
			} else if opCode == 1 { // bxl
				rb = rb ^ operand
			} else if opCode == 2 { // bst
				rb = combo(operand) % 8
			} else if opCode == 3 { // jnz
				if ra == 0 {
					continue
				}
				iptr = operand
			} else if opCode == 4 { // bxc
				rb = rb ^ rc
			} else if opCode == 5 { // out
				v := combo(operand) % 8
				// fmt.Print(strconv.Itoa(v) + ",")
				output += strconv.Itoa(v) + ","
			} else if opCode == 6 { // bdv
				rb = ra / int(math.Pow(2, float64(combo(operand))))
			} else if opCode == 7 { // cdv
				rc = ra / int(math.Pow(2, float64(combo(operand))))
			}
		}
		// fmt.Println("TRYING:", strconv.Itoa(origA), output)
		// return output == (strings.Join(programs, ",") + ",")
		return output[:len(output)-1]
	}

	fmt.Println("Part1:", runProgram())

	// After printing out the first iterations 0-10000ish, it's pretty clear
	// that the outputs start at 1 digit and go up with each power of 8 that is
	// used. So, for the 16 element puzzle input, the answer *should* be
	// between 8^15 and 8^16.
	// I'm sure there's some wicked math that can be done, but I just did a
	// manual pattern searching starting from the last digit. Fortunately,
	// there was enough of a consistent behavior that searching manually was
	// feasible, natch not pleasant. This is my attempt and reconstructing
	// my monkey brain's approach.

	// Starting with the last digit, search for the appropriate range.
	skipFactor := 1

	pNoCommas := strings.Join(programs, "")

	start := int(math.Pow(8, float64(len(programs)-1))) // 8^15
	end := int(math.Pow(8, float64(len(programs))))     // 8^16
	res := ""

outer:
	for {
		newStart := 0
		newEnd := 0
		skip := int(math.Pow10(len(programs) - (5 + skipFactor)))
		inZone := false
		lastSearch := start
		checkFrom := len(programs) - skipFactor

		// fmt.Println("LOOP:", start, end, skip)
		for i := start; i < end; i += skip {
			ra = i
			rb = 0
			rc = 0
			res = runProgram()
			rNoCommas := strings.ReplaceAll(res, ",", "")

			if rNoCommas == pNoCommas {
				fmt.Println("Part2:", i)
				break outer
			}

			// fmt.Println("MATCH CHECK:", rNoCommas[checkFrom:], pNoCommas[checkFrom:])
			if !inZone && rNoCommas[checkFrom:] == pNoCommas[checkFrom:] {
				// fmt.Println("MATCH BEG:", i, lastSearch, rNoCommas, rNoCommas[checkFrom:], pNoCommas[checkFrom:])
				newStart = lastSearch
				inZone = true
			} else if inZone && rNoCommas[checkFrom:] != pNoCommas[checkFrom:] {
				// fmt.Println("MATCH END:", i, lastSearch, rNoCommas, rNoCommas[checkFrom:], pNoCommas[checkFrom:])
				newEnd = i
				inZone = false
				break
			}
			lastSearch = i
		}

		// fmt.Println("RANGE:", newStart, newEnd, res)
		start = newStart
		end = newEnd
		skipFactor += 1
	}
}
