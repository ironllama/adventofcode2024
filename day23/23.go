package main

import (
	"aoc2024/utils"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(file), "\n")
	// fmt.Println("LINES:", lines)

	// Process the input, with each or pair connected as a key to the other of the pair
	// ie. "key: [neighbor]" with "co-de" into "co: [de]" and "de: [co]"
	links := map[string][]string{}
	for _, l := range lines {
		pair := strings.Split(l, "-")
		if _, ok := links[pair[0]]; ok {
			links[pair[0]] = append(links[pair[0]], pair[1])
		} else {
			links[pair[0]] = []string{pair[1]}
		}
		if _, ok := links[pair[1]]; ok {
			links[pair[1]] = append(links[pair[1]], pair[0])
		} else {
			links[pair[1]] = []string{pair[0]}
		}
	}
	// fmt.Println("MAP:")
	// for k, v := range links {
	// 	fmt.Println(k, ":", v)
	// }

	part1 := func() {
		// Find all combos of 3 connected to each other using the map above
		three := []string{}
		for k, v := range links { // For each key
			for in1, n1 := range v { // For each neighbor
				for in2 := in1 + 1; in2 < len(v); in2++ { // Check each other neighbor
					n2 := v[in2]
					if slices.Contains(links[n2], n1) && slices.Contains(links[n2], k) {
						newGroup := []string{k, n1, n2}
						slices.Sort(newGroup)
						newKey := fmt.Sprintf("%s,%s,%s", newGroup[0], newGroup[1], newGroup[2])
						three = append(three, newKey) // Can check for dupes now, or just remove later
					}
				}
			}
		}

		// Will be dupes, since co-de-ka will come from de-co-ka and ka-de-co
		three = utils.SliceRemoveDuplicatesFrom(three)
		slices.Sort(three) // To match text

		// fmt.Println(three)
		p1Total := 0
		for _, t := range three {
			if string(t[0]) == "t" {
				p1Total += 1
				continue
			}
			if strings.Contains(t, ",t") {
				p1Total += 1
			}
		}
		fmt.Println("Part1:", p1Total)
	}
	part1()

	part2 := func() {
		// LOL! Wishful thinking. Burned a bit of time, but was a gamble in favor of not having to think about the problem.
		// The gamble did NOT pay off. Wasted a lot of time wondering why there were 40 possibilities of length 12, but 0 of length 13.
		// Maybe there's a subtle bug in the code and this approach would work? Dunno. Anyway, at least worth a chuckle, in retrospect.
		// allCombos := []string{}
		// for k, v := range links { // For each key
		// 	for iv2, v2 := range v { // For each neighbor
		// 		for iv3 := iv2 + 1; iv3 < len(v); iv3++ { // Check each other neighbor
		// 			v3 := v[iv3]
		// 			for iv4 := iv3 + 1; iv4 < len(v); iv4++ { // Check each other neighbor
		// 				v4 := v[iv4]
		// 				for iv5 := iv4 + 1; iv5 < len(v); iv5++ { // Check each other neighbor
		// 					v5 := v[iv5]
		// 					for iv6 := iv5 + 1; iv6 < len(v); iv6++ { // Check each other neighbor
		// 						v6 := v[iv6]
		// 						for iv7 := iv6 + 1; iv7 < len(v); iv7++ { // Check each other neighbor
		// 							v7 := v[iv7]
		// 							for iv8 := iv7 + 1; iv8 < len(v); iv8++ { // Check each other neighbor
		// 								v8 := v[iv8]
		// 								for iv9 := iv8 + 1; iv9 < len(v); iv9++ { // Check each other neighbor
		// 									v9 := v[iv9]
		// 									for iv10 := iv9 + 1; iv10 < len(v); iv10++ { // Check each other neighbor
		// 										v10 := v[iv10]
		// 										for iv11 := iv10 + 1; iv11 < len(v); iv11++ { // Check each other neighbor
		// 											v11 := v[iv11]
		// 											for iv12 := iv11 + 1; iv12 < len(v); iv12++ { // Check each other neighbor
		// 												v12 := v[iv12]
		// 												for iv13 := iv12 + 1; iv13 < len(v); iv13++ { // Check each other neighbor
		// 													v13 := v[iv13]
		// 													// 		for iv14 := iv3 + 1; iv14 < len(v); iv14++ { // Check each other neighbor
		// 													// 			v8 := v[iv14]
		// 													// 			for iv9 := iv8 + 1; iv9 < len(v); iv9++ { // Check each other neighbor
		// 													// 				v9 := v[iv9]
		// 													// for ivlast := iv12 + 1; ivlast < len(v); ivlast++ { // Check each other neighbor
		// 													// 	vLast := v[ivlast]
		// 													for ivlast := iv13 + 1; ivlast < len(v); ivlast++ { // Check each other neighbor
		// 														vLast := v[ivlast]
		// 														if slices.Contains(links[vLast], k) &&
		// 															slices.Contains(links[vLast], v2) &&
		// 															slices.Contains(links[vLast], v3) &&
		// 															slices.Contains(links[vLast], v4) &&
		// 															slices.Contains(links[vLast], v5) &&
		// 															slices.Contains(links[vLast], v6) &&
		// 															slices.Contains(links[vLast], v7) &&
		// 															slices.Contains(links[vLast], v8) &&
		// 															slices.Contains(links[vLast], v9) &&
		// 															slices.Contains(links[vLast], v10) &&
		// 															slices.Contains(links[vLast], v11) &&
		// 															// slices.Contains(links[vLast], v12) {
		// 															slices.Contains(links[vLast], v12) &&
		// 															slices.Contains(links[vLast], v13) {
		// 															// newGroup := []string{k, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, vLast}
		// 															newGroup := []string{k, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, vLast}
		// 															slices.Sort(newGroup)
		// 															// fmt.Println("FOUND:", newGroup)
		// 															newKey := fmt.Sprintf(
		// 																// "%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		// 																"%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		// 																newGroup[0],
		// 																newGroup[1],
		// 																newGroup[2],
		// 																newGroup[3],
		// 																newGroup[4],
		// 																newGroup[5],
		// 																newGroup[6],
		// 																newGroup[7],
		// 																newGroup[8],
		// 																newGroup[9],
		// 																newGroup[10],
		// 																newGroup[11],
		// 																// newGroup[12])
		// 																newGroup[12],
		// 																newGroup[13])
		// 															allCombos = append(three, newKey)
		// 														}
		// 														// 		}
		// 														// 	}
		// 													}
		// 												}
		// 											}
		// 										}
		// 									}
		// 								}
		// 							}
		// 						}
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		allLongest := []string{}
		for k, v := range links { // For each key
			pn := make([]string, len(v)+1) // Create a copy of list of neighbors, but including self
			copy(pn, v)
			pn[len(v)-1] = k

			// Generate a list of neighbors that are "strongly connected", using key and first neighbor
			longest := []string{}
			for _, n1 := range v { // For each neighbor
				cn := links[n1]  // List of neighbors neighbors
				nm := []string{} // List of matches on parent (n1) neighbor
				nm = append(nm, n1)

				for _, tpn := range pn { // Compare parent neighbor list with neighbor neighbor list
					if slices.Contains(cn, tpn) { // At this point, I'm thinking, I need to use longer var names
						nm = append(nm, tpn)
					}
				}

				// fmt.Println("FOR:", k, n1, "BIG:", nm)
				if len(nm) > len(longest) {
					longest = nm
				}
			}
			slices.Sort(longest) // For easier rollup, later
			allLongest = append(allLongest, strings.Join(longest, ","))
		}

		// Find the longest, which would have the same list the same number of times.
		// For example, a network of 5 should appear 5 times, each generated by one of the 5
		// To make it easier to find, break out to a frequency map and then make sure the numbers match.
		allLongestOccurs := map[string]int{}
		for _, al := range allLongest {
			if _, ok := allLongestOccurs[al]; ok {
				allLongestOccurs[al] += 1
			} else {
				allLongestOccurs[al] = 1
			}
		}

		longest := ""
		longestSize := 0
		for k, v := range allLongestOccurs {
			toks := strings.Split(k, ",")
			if len(toks) == v && (longestSize == 0 || v > longestSize) {
				longest = k
				longestSize = v
			}
		}

		fmt.Println("Part2:", longest)
	}
	part2()
}
