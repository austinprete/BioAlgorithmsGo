package main

import (
	"fmt"

	"github.com/austinprete/bio_algorithms/utilities"
)

func findPatternMatches(pattern string, genome string) []int {

	matchIndices := []int{}

	for index := 0; index <= len(genome)-len(pattern); index++ {
		substring := genome[index : index+len(pattern)]

		if substring == pattern {
			matchIndices = append(matchIndices, index)
		}
	}

	return matchIndices
}

func main() {
	lines := utilities.ReadLinesFromDataset("ba1d.txt")

	pattern := lines[0]
	genome := lines[1]

	result := findPatternMatches(pattern, genome)

	for _, matchIndex := range result {
		fmt.Printf("%d ", matchIndex)
	}
}
