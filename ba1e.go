package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/austinprete/bio_algorithms/utilities"
)

func findClumps(genome string, k int, L int, t int) []string {
	defer utilities.TimeTrack(time.Now(), "findClumps")

	clumpedPatternMap := make(map[string]bool)

	text := genome[:L]

	frequencyArray := utilities.ComputeFrequencies(text, k)

	for index := 0; index < len(frequencyArray); index++ {
		if frequencyArray[index] >= t {
			pattern := utilities.NumberToPattern(index, k)
			clumpedPatternMap[pattern] = true
		}
	}

	for index := 1; index <= len(genome)-L; index++ {
		firstPattern := genome[index-1 : (index+k)-1]

		patternIndex := utilities.PatternToNumber(firstPattern)
		frequencyArray[patternIndex] -= 1

		lastPattern := genome[(index+L)-k : index+L]

		patternIndex = utilities.PatternToNumber(lastPattern)
		frequencyArray[patternIndex] += 1

		if frequencyArray[patternIndex] >= t {
			clumpedPatternMap[lastPattern] = true
		}
	}

	clumpedPatternArray := []string{}

	for pattern := range clumpedPatternMap {
		clumpedPatternArray = append(clumpedPatternArray, pattern)
	}

	return clumpedPatternArray
}

func main() {
	lines := utilities.ReadLinesFromDataset("ba1e.txt")

	genome := lines[0]
	splitIntegerInputs := strings.Split(lines[1], " ")

	k, _ := strconv.Atoi(splitIntegerInputs[0])
	L, _ := strconv.Atoi(splitIntegerInputs[1])
	t, _ := strconv.Atoi(splitIntegerInputs[2])

	result := findClumps(genome, k, L, t)

	for _, pattern := range result {
		fmt.Printf("%s ", pattern)
	}
	fmt.Println()

}

//// NOTE: This is a bad implementation. The use of go-routines allows it to finish in a reasonable amount of time
//// on the Rosalind datasets, but an improved algorithm is needed.
//func findClumpsConcurrent(genome string, k int, L int, t int) []string {
//	kmers := make(map[string]bool)
//	clumpedKmers := make(map[string]bool)
//
//	for index := 0; index <= len(genome)-k; index++ {
//		substring := genome[index : index+k]
//		kmers[substring] = true
//	}
//
//	for index := 0; index <= len(genome)-L; index++ {
//		substring := genome[index : index+L]
//
//		ch := make(chan struct {
//			string
//			bool
//		}, len(kmers))
//
//		for kmer := range kmers {
//
//			go func(substring, kmer string) {
//
//				count := utilities.PatternCount(substring, kmer)
//
//				ch <- struct {
//					string
//					bool
//				}{kmer, count >= t}
//
//			}(substring, kmer)
//
//		}
//
//		for index := 0; index < len(kmers); index++ {
//			pair := <-ch
//
//			if pair.bool {
//				clumpedKmers[pair.string] = true
//			}
//		}
//	}
//
//	clumpedKmersList := []string{}
//	for kmer := range clumpedKmers {
//		clumpedKmersList = append(clumpedKmersList, kmer)
//	}
//
//	return clumpedKmersList
//}
