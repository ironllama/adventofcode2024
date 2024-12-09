package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	type block struct {
		val  int
		size int // For part2
	}

	blocks := []block{}
	currNum := 0
	doSpace := false
	for _, c := range input { // For each char in the disk map
		num, _ := strconv.Atoi(string(c))
		for k := 0; k < num; k++ { // Use the num to control the num of chars appearing
			if doSpace {
				// blocks = append(blocks, block{-1})  // part1 only
				blocks = append(blocks, block{-1, num}) // part2 (backwards compatible with Part1)
			} else {
				// blocks = append(blocks, block{currNum})  // part1 only
				blocks = append(blocks, block{currNum, num}) // part2 (backwards compatible with Part2)
			}
		}
		if !doSpace {
			currNum += 1
		}
		doSpace = !doSpace
	}
	// log.Println("BLOCKS:", blocks)

	printBlocks := func(pBlocks []block) {
		fmt.Print("BLOCKS: ") // The log.Print still adds a newline. Annoying.
		for _, b := range pBlocks {
			if b.val == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(b.val)
			}
		}
		fmt.Println("")
	}
	_ = printBlocks // Ugly way of telling Go to ignore that I may never use this name

	totalBlocks := func(tBlocks []block) int {
		total := 0
		for i, b := range tBlocks {
			if b.val != -1 {
				total += b.val * i
			}
		}
		return total
	}

	part1 := func() {
		p1Blocks := make([]block, len(blocks))
		copy(p1Blocks, blocks)

		for {
			nextSpace := slices.IndexFunc(p1Blocks, func(b block) bool { return b.val == -1 })
			var nextBlockPos int
			for i := len(p1Blocks) - 1; i >= 0; i-- {
				if p1Blocks[i].val != -1 {
					nextBlockPos = i
					break
				}
			}
			if nextSpace < nextBlockPos {
				p1Blocks[nextSpace], p1Blocks[nextBlockPos] = p1Blocks[nextBlockPos], p1Blocks[nextSpace]
			} else {
				break
			}
		}
		// printBlocks(p1Blocks)

		log.Println("Part1:", totalBlocks(p1Blocks))
	}
	part1()

	part2 := func() {
		findBlockPos := len(blocks) - 1                                                     // Assuming the string never ends with .'s
		findSpacePos := slices.IndexFunc(blocks, func(b block) bool { return b.val == -1 }) // First open space

		nextVal := blocks[findBlockPos].val

		for nextVal >= 0 {
			// log.Println("START:", findBlockPos, findSpacePos, "VAL:", blocks[findBlockPos])
			currBlock := blocks[findBlockPos]
			findSpacePos = 0 // Reset open space search from beginning.

			// Find the appropriate open space
			found := false
		outer:
			for i := findSpacePos; i < len(blocks)-1; i++ {
				if blocks[i].val == -1 {
					for k := 0; k < currBlock.size; k++ {
						if (i+k) > len(blocks)-1 || blocks[i+k].val != -1 {
							continue outer
						}
					}
					found = true
					findSpacePos = i
					break outer
				}
			}
			// log.Println("FOUND:", found, findSpacePos, "FOR:", currBlock.val)

			// Swap or skip if no match
			if found && findSpacePos < findBlockPos {
				for i := 0; i < currBlock.size; i++ {
					blocks[findBlockPos-i], blocks[findSpacePos+i] = blocks[findSpacePos+i], blocks[findBlockPos-i]
				}
			}
			// printBlocks(blocks)

			// Find next interesting block
			nextVal -= 1
			findBlockPos -= currBlock.size       // Move over the size of the block. NOTE: after swap, need to use old space pos.
			for i := findBlockPos; i >= 0; i-- { // Scan for next value, if -1, keep scanning
				if blocks[i].val != -1 && blocks[i].val == nextVal {
					findBlockPos = i
					break
				}
			}
			// log.Println("NEW POS:", findBlockPos, findSpacePos)
		}

		log.Println("Part2:", totalBlocks(blocks))
	}
	part2()
}
