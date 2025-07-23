package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// data, _ := os.ReadFile("cmd/aoc2018/day3/sample.txt")
	data, _ := os.ReadFile("cmd/aoc2018/day3/input.txt")

	lines := strings.Split(string(data), "\n")

	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func partOne(lines []string) int {

	totalArea := 0
	var fabric [1000][1000]int

	for _, value := range lines {
		var id, x, y, width, height int
		_, err := fmt.Sscanf(value, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)

		if err != nil {
			continue
		}

		for row := x; row < x+width; row++ {
			for column := y; column < y+height; column++ {
				fabric[row][column]++
				if fabric[row][column] == 2 {
					totalArea++
				}
			}
		}
	}

	return totalArea

}

type Claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

func partTwo(lines []string) int {

	valid := make(map[int]bool)

	var claims []Claim

	for _, value := range lines {
		var id, x, y, width, height int
		_, err := fmt.Sscanf(value, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)

		if err != nil {
			continue
		}

		claims = append(claims, Claim{ID: id, X: x, Y: y, Width: width, Height: height})
	}

	for _, c := range claims {
		valid[c.ID] = true
	}

	for i := 0; i < len(claims); i++ {
		for j := 0; j < len(claims); j++ {
			a := claims[i]
			b := claims[j]

			if a.ID == b.ID {
				continue
			}

			if overlaps(a, b) {
				valid[a.ID] = false
				valid[b.ID] = false
			}
		}
	}

	fmt.Println(valid)
	var ID int
	for key, value := range valid {
		if value {
			ID = key
		}
	}

	return ID
}

func overlaps(a, b Claim) bool {
	return !(a.X+a.Width <= b.X || b.X+b.Width <= a.X || a.Y+a.Height <= b.Y || b.Y+b.Height <= a.Y)

}
