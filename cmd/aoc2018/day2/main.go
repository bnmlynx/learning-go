package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	data, _ := os.ReadFile("cmd/aoc2018/day2/input.txt")

	lines := strings.Split(string(data), "\n")

	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

// abbccc

func partOne(lines []string) int {
	// take each line and does can it match
	twoLetterCount := 0
	threeLetterCount := 0
	for _, line := range lines {
		letters := strings.Split(string(line), "")
		slices.Sort(letters)
		var previousLetter string
		letterCount := 1
		twoLetterFound := false
		threeLetterFound := false
		for k, letter := range letters {
			if letter == previousLetter {
				letterCount++
			}
			if letter != previousLetter || k+1 == len(letters) {
				if letterCount == 2 && !twoLetterFound {
					twoLetterCount++
					twoLetterFound = true
				}

				if letterCount == 3 && !threeLetterFound {
					threeLetterCount++
					threeLetterFound = true
				}
				// reset the letter count
				letterCount = 1
			}

			previousLetter = letter
		}
	}

	return twoLetterCount * threeLetterCount
}

func partTwo(lines []string) string {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {

			diffCount := 0
			diffIndex := -1
			for k := 0; k < len(lines[i]); k++ {
				if lines[i][k] != lines[j][k] {
					diffCount++
					diffIndex = k

					if diffCount > 1 {
						break
					}
				}

			}
			if diffCount == 1 {
				return lines[i][:diffIndex] + lines[j][diffIndex+1:]
			}
		}
	}

	return ""
}
