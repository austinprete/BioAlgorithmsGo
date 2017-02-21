package main

import (
	"fmt"
	"strconv"

	"github.com/austinprete/bio_algorithms/utilities"
)

func findMostFrequentKmer(text string, k int) []string {
	kmersAndCounts := make(map[string]int)

	for index := 0; index <= len(text)-k; index++ {
		kmer := text[index : index+k]

		_, kmerPresent := kmersAndCounts[kmer]
		if !kmerPresent {
			kmersAndCounts[kmer] = 1
		} else {
			kmersAndCounts[kmer] += 1
		}
	}

	highestCount := 0

	mostFrequentKmers := []string{}

	for kmer, count := range kmersAndCounts {
		if count > highestCount {
			highestCount = count
			mostFrequentKmers = []string{kmer}
		} else if count == highestCount {
			mostFrequentKmers = append(mostFrequentKmers, kmer)
		}
	}

	return mostFrequentKmers
}

func main() {
	lines := utilities.ReadLinesFromDataset("ba1b.txt")

	text := lines[0]
	k, err := strconv.Atoi(lines[1])

	if err != nil {
		panic(err)
	}

	result := findMostFrequentKmer(text, k)

	for _, kmer := range result {
		fmt.Printf("%s ", kmer)
	}
}
