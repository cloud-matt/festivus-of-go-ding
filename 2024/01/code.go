package main

import (
	"cmp"
	"github.com/jpillora/puzzler/harness/aoc"
	"slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	// Init
	lines := strings.Split(strings.TrimSpace(input), "\n")
	groupANums := make([]int, 0, len(lines))
	groupBNums := make([]int, 0, len(lines))

	// Populate groupANums and groupBNums
	for _, line := range lines {
		pair := strings.SplitN(line, " ", 2)
		if part2 {
			// If part 2, we don't need to sort
			groupANums = append(groupANums, ConvertToNum(pair[0]))
			groupBNums = append(groupBNums, ConvertToNum(pair[1]))
		} else {
			// If part 1, we DO need to sort
			groupANums = InsertSorted(groupANums, ConvertToNum(pair[0]))
			groupBNums = InsertSorted(groupBNums, ConvertToNum(pair[1]))
		}
	}

	// Calculate the answer for the requested part
	if part2 {
		return CalculateSimilarityScore(groupANums, groupBNums)
	} else {
		return CalculateDistance(groupANums, groupBNums)
	}
}

func CalculateSimilarityScore(groupANums []int, groupBNums []int) int {
	// Build frequency map for group B
	groupBFreqMap := make(map[int]int)
	for _, num := range groupBNums {
		groupBFreqMap[num] = groupBFreqMap[num] + 1
	}

	// Calculate similarity score from group A
	similarityScore := 0
	for _, num := range groupANums {
		if groupBFreqMap[num] >= 1 {
			similarityScore += num * groupBFreqMap[num]
		}
	}

	return similarityScore
}

func CalculateDistance(groupANums []int, groupBNums []int) int {
	distance := 0
	for index := range groupANums {
		distance += GetAbsDifference(groupANums[index], groupBNums[index])
	}
	return distance
}

// ConvertToNum Converts string to number
func ConvertToNum(numString string) int {
	numInt, err := strconv.Atoi(strings.TrimSpace(numString))
	if err != nil {
		panic(err)
	}
	return numInt
}

// InsertSorted inserts value t into array ts while maintaining a sorted ts
// T is required to be a cmp.Ordered type
func InsertSorted[T cmp.Ordered](ts []T, t T) []T {
	// Search for the slot to insert t
	i, _ := slices.BinarySearch(ts, t)

	// Make room for new value
	ts = append(ts, *new(T))
	// Copy the window over one space to the right
	copy(ts[i+1:], ts[i:])
	// Add in the new value
	ts[i] = t
	return ts
}

// GetAbsDifference returns the absolute difference between a and b
// This avoids the need to convert to/from float when using math.Abs
func GetAbsDifference(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
