package utilities

import (
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLinesFromDataset(fileName string) []string {
	file_data, err := ioutil.ReadFile("resources/" + fileName)
	check(err)

	contents := string(file_data)

	lines := strings.Split(contents, "\n")

	return lines
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
