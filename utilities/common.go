package utilities

import "math"

func ComputeFrequencies(text string, k int) (frequencyArray []int) {

	for index := 0; index < int(math.Pow(4, float64(k))); index++ {
		frequencyArray = append(frequencyArray, 0)
	}

	for index := 0; index <= len(text)-k; index++ {
		pattern := text[index : index+k]

		patternIndex := PatternToNumber(pattern)
		frequencyArray[patternIndex] += 1
	}

	return frequencyArray
}

func PatternCount(text string, pattern string) int {
	count := 0
	for index := 0; index <= len(text)-len(pattern); index++ {
		if text[index:index+len(pattern)] == pattern {
			count++
		}
	}

	return count
}

func NumberToPattern(index, k int) string {
	NUMBER_TO_LETTER := map[int]string{0: "A", 1: "C", 2: "G", 3: "T"}

	if k == 1 {
		return NUMBER_TO_LETTER[index]
	}

	prefixIndex := int(index / 4)
	remainder := index % 4

	letter := NUMBER_TO_LETTER[remainder]

	prefixPattern := NumberToPattern(prefixIndex, k-1)

	return prefixPattern + letter
}

func PatternToNumber(pattern string) int {
	LETTER_TO_NUMBER := map[string]int{"A": 0, "C": 1, "G": 2, "T": 3}

	patternLength := len(pattern)

	if patternLength == 0 {
		return 0
	}

	prefix := pattern[:patternLength-1]
	lastSymbol := string(pattern[patternLength-1])

	return (4 * PatternToNumber(prefix)) + LETTER_TO_NUMBER[lastSymbol]
}
