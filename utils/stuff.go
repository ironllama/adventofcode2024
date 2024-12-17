package utils

import (
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// ============================================================================
// Grid functions for 2D maps in Y,X format. Grid itself contains string chars.
var GridYXKey string = "%d,%d"
var GridNeighbors [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // YX notation: N E S W or U R D L
var GridDiags [][]int = [][]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}   // YX notation: NE SE SW NW
var GridAllAround = append(GridNeighbors, GridDiags...)

func IsValidPosInGrid(y int, x int, grid [][]string) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}

// ============================================================================
// Slice functions.
func SliceContains(slice []string, element string) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}

func SliceSum(array []int) int {
	sum := 0
	for _, item := range array {
		sum += item
	}
	return sum
}

func SliceReverse(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		pos := len(a) - 1 - i
		a[i], a[pos] = a[pos], a[i]
	}
	return a
}

func SliceRemoveDuplicatesFrom(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// ============================================================================
// Set functions for a Set of string values. Using struct{} for minimal value storage cost.
// NOTE: These are implemented as a map and are therefore UNORDERED.
// Source: https://www.bytesizego.com/blog/set-in-golang
type Set struct {
	elements map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[string]struct{}),
	}
}

func (s *Set) Add(value string) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set
func (s *Set) Remove(value string) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set
func (s *Set) Contains(value string) bool {
	_, found := s.elements[value]
	return found
}

// Size returns the number of elements in the set
func (s *Set) Size() int {
	return len(s.elements)
}

// List returns all elements in the set as a slice
func (s *Set) List() []string {
	keys := make([]string, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func (s *Set) Union(other *Set) *Set {
	result := NewSet()
	for key := range s.elements {
		result.Add(key)
	}
	for key := range other.elements {
		result.Add(key)
	}
	return result
}

func (s *Set) Intersection(other *Set) *Set {
	result := NewSet()
	for key := range s.elements {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

func (s *Set) Difference(other *Set) *Set {
	result := NewSet()
	for key := range s.elements {
		if !other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}
