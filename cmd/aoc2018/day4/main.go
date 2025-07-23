package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	data, _ := os.ReadFile("cmd/aoc2018/day4/sample.txt")

	lines := strings.Split(string(data), "\n")

	fmt.Println(partOne(lines))
	// fmt.Println(partTwo(lines))
}

func partOne(lines []string) int {

	orderChronologically(lines)

	return 0
}

func orderChronologically(lines []string) {

	for _, value := range lines {

		r := regexp.MustCompile(`\[(.+)\].*#(\d+)`)
		// r := regexp.MustCompile(`\[(.+)\]+(.+)`)

		match := r.FindStringSubmatch(value)

		// fmt.Println(value)
		fmt.Println(match)

	}

}

type Test struct {
	ID     int
	Asleep time.Time
	Awake  time.Time
}
