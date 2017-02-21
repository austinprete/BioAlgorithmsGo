package main

import (
	"fmt"

	"github.com/austinprete/bio_algorithms/utilities"
)

func main() {
	lines := utilities.ReadLinesFromDataset("ba1a.txt")

	text := lines[0]
	pattern := lines[1]

	result := utilities.PatternCount(text, pattern)

	fmt.Println(result)
}
