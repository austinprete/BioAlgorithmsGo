package main

import (
	"bytes"
	"fmt"

	"github.com/austinprete/bio_algorithms/utilities"
)

var nucleotideMappings map[string]string = map[string]string{"A": "T", "C": "G", "G": "C", "T": "A"}

func findReverseComplement(pattern string) string {
	reversedPattern := utilities.Reverse(pattern)

	var complement bytes.Buffer

	for index := range reversedPattern {
		newLetter := nucleotideMappings[string(reversedPattern[index])]

		complement.WriteString(newLetter)
	}

	return complement.String()
}

func main() {
	lines := utilities.ReadLinesFromDataset("ba1c.txt")

	pattern := lines[0]

	result := findReverseComplement(pattern)

	fmt.Println(result)
}
