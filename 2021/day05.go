package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day5() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	coordinates, xn, yn := parseInput(input)
	fmt.Println("Bounds are: ", xn, yn)
	matrix := markLineSegments(coordinates, xn, yn, false)
	score := calculateScore(matrix, 2)
	logResult(5, 1, "Overlap count is: ", score)
	matrix = markLineSegments(coordinates, xn, yn, true)
	score = calculateScore(matrix, 2)
	logResult(5, 2, "Overlap count (with diagonals) is: ", score)
	return nil
}

type Coordinates struct {
	x1, y1, x2, y2 int
}

func calculateScore(matrix [][]int, threshold int) int {
	var res int
	for _, row := range matrix {
		for _, p := range row {
			if p >= threshold {
				res++
			}
		}
	}
	return res
}

func markLineSegments(input []Coordinates, xn, yn int, includeDiagonal bool) [][]int {
	// make nxn matrix instead of mxn
	var xyn int
	if xn > yn {
		xyn = xn
	} else {
		xyn = yn
	}
	matrix := make([][]int, xyn+1)
	for i := range matrix {
		matrix[i] = make([]int, xyn+1)
	}
	for _, c := range input {
		if c.x1 == c.x2 {
			// horizontal
			var y1, y2 int
			if c.y1 > c.y2 {
				y1 = c.y2
				y2 = c.y1
			} else {
				y1 = c.y1
				y2 = c.y2
			}
			for i := y1; i <= y2; i++ {
				matrix[c.x1][i] += 1
			}
		} else if c.y1 == c.y2 {
			// vertical
			var x1, x2 int
			if c.x1 > c.x2 {
				x1 = c.x2
				x2 = c.x1
			} else {
				x1 = c.x1
				x2 = c.x2
			}
			for i := x1; i <= x2; i++ {
				matrix[i][c.y1] += 1
			}
		} else {
			// diagonal
			if includeDiagonal == false {
				continue
			}
			var xinc, yinc int
			// determine increments for loop. We could have 4 different angles of diagonals
			if c.x1 < c.x2 {
				xinc = 1
			} else {
				xinc = -1
			}
			if c.y1 < c.y2 {
				yinc = 1
			} else {
				yinc = -1
			}
			for i, j := c.x1, c.y1; ; i, j = i+xinc, j+yinc {
				matrix[i][j] += 1
				// both the coorindates are inclusive. So count first, check afterwards
				if i == c.x2 || j == c.y2 {
					//fmt.Println("Coordinates: ", c, " i ", i, " j ", j)
					break
				}
			}
		}
	}
	return matrix
}

func parseInput(input []string) ([]Coordinates, int, int) {
	var xn, yn int // farthest points in 2-d map
	var res []Coordinates
	for _, l := range input {
		// strip " -> "
		parts := strings.Split(l, " -> ")
		if len(parts) != 2 {
			fmt.Println("Failed to parse line: ", l)
		}
		// extract first coordinate
		c1 := strings.Split(parts[0], ",")
		// extract second coordinate
		c2 := strings.Split(parts[1], ",")
		// convert it to integer
		var x1, y1, x2, y2 int
		if n, err := strconv.Atoi(c1[0]); err != nil {
			fmt.Println("Failed parsing coordinate: ", err)
		} else {
			x1 = n
		}
		if n, err := strconv.Atoi(c1[1]); err != nil {
			fmt.Println("Failed parsing coordinate: ", err)
		} else {
			y1 = n
		}
		if n, err := strconv.Atoi(c2[0]); err != nil {
			fmt.Println("Failed parsing coordinate: ", err)
		} else {
			x2 = n
		}
		if n, err := strconv.Atoi(c2[1]); err != nil {
			fmt.Println("Failed parsing coordinate: ", err)
		} else {
			y2 = n
		}
		// populate coordinates type
		c := Coordinates{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
		res = append(res, c)
		// Calculate bounds of the matrix
		if x1 > xn {
			xn = x1
		}
		if x2 > xn {
			xn = x2
		}
		if y1 > yn {
			yn = y1
		}
		if y2 > yn {
			yn = y2
		}
	}
	return res, xn, yn
}
