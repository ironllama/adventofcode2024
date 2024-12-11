package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	toks := strings.Fields(string(input))
	// log.Println("TOKS:", toks)

	part1 := func() {
		blink := func(inToks []string) []string {
			newToks := []string{}

			for _, tok := range inToks {
				if tok == "0" {
					newToks = append(newToks, "1")
					continue
				}

				if len(tok)%2 == 0 {
					halfIdx := len(tok) / 2 // Already floor division, as ints

					newToks = append(newToks, tok[:halfIdx])

					secondPart := tok[halfIdx:]
					secondPart = strings.TrimLeft(secondPart, "0")
					if len(secondPart) == 0 {
						secondPart = "0"
					}
					newToks = append(newToks, secondPart)
					continue
				}

				num, _ := strconv.Atoi(tok)
				newToks = append(newToks, strconv.Itoa(num*2024))
			}
			return newToks
		}

		for range 25 {
			toks = blink(toks)
			// log.Println("TOKS:", toks)
			// fmt.Println("TOKS:", toks)
		}
		log.Println("Part1:", len(toks))
	}
	part1()

	part2 := func() {
		history := make(map[string]int)

		var blink func(tok string, num int) int // Pre-declared so we can use it to recurse

		checkHistory := func(val string, num int) int {
			key := val + "," + strconv.Itoa(num)
			newNum, ok := history[key]
			if !ok {
				newNum = blink(val, num)
				history[key] = newNum
			}
			return newNum
		}

		limit := 75
		count := 0
		blink = func(tok string, num int) int {
			// log.Println("blink:", tok, num)
			num += 1

			if num == (limit + 1) {
				return 1
			}

			if tok == "0" {
				return checkHistory("1", num)
			}

			if len(tok)%2 == 0 {
				halfIdx := len(tok) / 2 // Already floor division, as ints

				checkSum := checkHistory(tok[:halfIdx], num)

				secondPart := tok[halfIdx:]
				secondPart = strings.TrimLeft(secondPart, "0")
				if len(secondPart) == 0 {
					secondPart = "0"
				}
				checkSum += checkHistory(secondPart, num)
				return checkSum
			}

			miscNum, _ := strconv.Atoi(tok)
			return checkHistory(strconv.Itoa(miscNum*2024), num)
		}

		for _, tok := range strings.Fields(string(input)) {
			count += blink(tok, 0)
		}
		log.Println("Part2:", count)
	}
	part2()
}
