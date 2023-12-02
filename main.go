package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v\n", err)
	}

	lines := readLines(f)
	linesCopy := readLines(f)

	for i := range lines {
		fmt.Printf("origLine : '%v'\n", linesCopy[i])
		out := replaceNumberWordsInOrder(lines[i])
		fmt.Printf("new line :")
		lines[i] = out
	}

	for _, line := range lines {
		fmt.Printf("line : '%v'\n", line)
	}

	var sum int
	for _, line := range lines {
		first, last := findFirstAndLastNumbers(line)
		fmt.Printf("line: '%v'\n", line)
		fmt.Printf("first, last: '%v'-'%v'\n", first, last)
		sum += intConvertFatal(fmt.Sprintf("%v%v", first, last))
		// fmt.Printf("runningSum: '%v'\n", sum)
	}
	fmt.Printf("final sum: %v\n", sum)
}

func findFirstAndLastNumbers(line string) (int, int) {
	var first, last int
	firstFound, lastFound := false, false

	for _, lw := range lineWords {
		if strings.Contains(line, lw.word) {
			if !firstFound {
				first = lw.num
				firstFound = true
			}
			last = lw.num
			lastFound = true
		}
	}

	for _, rune := range line {
		if isDigit(byte(rune)) {
			num := int(rune - '0')
			if !firstFound {
				first = num
				firstFound = true
			}
			last = num
			lastFound = true
		}
	}

	if !lastFound {
		last = first
	}

	return first, last
}

func readLines(f []byte) []string {
	var lines []string
	var buffer string
	for _, char := range f {
		if char != '\n' && char != '\r' {
			buffer += string(char)
		} else {
			if buffer != "" {
				lines = append(lines, buffer)
			}
			buffer = ""
		}
	}
	if buffer != "" {
		lines = append(lines, buffer)
	}
	return lines
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func replaceNumberWordsInOrder(line string) string {
	type occurrence struct {
		index  int
		num    string
		length int
	}

	var occurrences []occurrence
	for _, wordInt := range lineWords {
		index := 0
		for {
			foundIndex := strings.Index(line[index:], wordInt.word)
			if foundIndex == -1 {
				break
			}
			actualIndex := index + foundIndex
			occurrences = append(occurrences,
				occurrence{
					actualIndex,
					strconv.Itoa(wordInt.num),
					len(wordInt.word),
				})
			index = actualIndex + 1 // Move past this occurrence
		}
	}

	sort.Slice(occurrences, func(i, j int) bool {
		return occurrences[i].index < occurrences[j].index
	})

	var newLine strings.Builder
	lastIndex := 0
	for _, occ := range occurrences {
		if occ.index >= lastIndex {
			newLine.WriteString(line[lastIndex:occ.index])
			newLine.WriteString(occ.num)
			lastIndex = occ.index + occ.length
		}
	}
	if lastIndex < len(line) {
		newLine.WriteString(line[lastIndex:])
	}

	return newLine.String()
}

func intConvertFatal(combinedStr string) int {
	combinedInt, err := strconv.Atoi(combinedStr)
	if err != nil {
		log.Fatalf("error converting string to int: %v", combinedStr)
	}
	return combinedInt
}

// var lineWords = map[string]int{
// 	"one":   1,
// 	"two":   2,
// 	"three": 3,
// 	"four":  4,
// 	"five":  5,
// 	"six":   6,
// 	"seven": 7,
// 	"eight": 8,
// 	"nine":  9,
// }

var lineWords = []struct {
	word string
	num  int
}{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
}
