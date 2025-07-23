package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("cmd/aoc2018/day1/input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Println("Part 1:", partOne(lines))
	fmt.Println("Part 2:", partTwo(lines))
}

func partOne(lines []string) int {
	total := 0
	for _, value := range lines {
		if value == "" {
			continue
		}
		num := 0
		num, _ = strconv.Atoi(value)
		total += num
	}
	return total
}

func partTwo(lines []string) int {
	// store all the resultant frequencies and if it happens twice break out and

	resultingFrequncies := make(map[int]bool)
	foundRepeating := true
	total := 0

	for foundRepeating {
		for _, value := range lines {
			if value == "" {
				continue
			}
			num := 0
			num, _ = strconv.Atoi(value)
			total = total + num

			if _, ok := resultingFrequncies[total]; !ok {
				resultingFrequncies[total] = true
			} else {
				// fmt.Println("Part 2444:", total)
				foundRepeating = false
				return total
			}

		}
	}

	return total
}
