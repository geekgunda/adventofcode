package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day13() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	count := countOrigamiDots(input)
	logResult(13, 1, "Count is: ", count)
	return nil
}

func countOrigamiDots(input []string) int {
	canvas, steps := parseOrigamiInput(input)
	var count int
	for i, step := range steps {
		parts := strings.Split(step, "=")
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing step instruction: ", err)
			return 0
		}
		switch parts[0] {
		case "x":
			canvas = foldVertically(canvas, val)
		case "y":
			canvas = foldHorizontally(canvas, val)
		default:
			fmt.Println("Invalid folding axis found: ", parts[0])
			return 0
		}
		if i == 0 {
			count = countRemainingDots(canvas)
		}
	}
	printCanvas(canvas)
	return count
}

func printCanvas(canvas [][]bool) {
	fmt.Println()
	for _, row := range canvas {
		for _, p := range row {
			if p {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func foldHorizontally(canvas [][]bool, y int) [][]bool {
	// Copy top half over directly
	foldedCanvas := canvas[:y]
	// Copy the remaining rows over starting from bottom
	currRow := y - 1
	for _, row := range canvas[y+1:] {
		for j, p := range row {
			if p {
				foldedCanvas[currRow][j] = p
			}
		}
		currRow--
	}
	return foldedCanvas
}

func foldVertically(canvas [][]bool, x int) [][]bool {
	foldedCanvas := make([][]bool, len(canvas))
	for i := 0; i < len(canvas); i++ {
		// copy left half over directly
		foldedCanvas[i] = canvas[i][:x]
	}
	for i, row := range canvas {
		// copy the right half starting from right and going over to left
		currCol := x - 1
		for _, p := range row[x+1:] {
			if p {
				foldedCanvas[i][currCol] = p
			}
			currCol--
		}
	}
	return foldedCanvas
}

func countRemainingDots(canvas [][]bool) int {
	var count int
	for _, row := range canvas {
		for _, p := range row {
			if p {
				count++
			}
		}
	}
	return count
}

type PaperDot struct {
	x, y int
}

func parseOrigamiInput(input []string) ([][]bool, []string) {
	var dots []PaperDot
	var steps []string
	var stepsBegan bool
	var xMax, yMax int
	for _, line := range input {
		if len(line) == 0 {
			stepsBegan = true
			continue
		}
		if stepsBegan == false {
			row, err := readLineAsNumbersCsv(line)
			if err != nil {
				fmt.Println("Error parsing dots: ", err)
			}
			dots = append(dots, PaperDot{x: row[0], y: row[1]})
			if row[0] > xMax {
				xMax = row[0]
			}
			if row[1] > yMax {
				yMax = row[1]
			}
		} else {
			step := strings.TrimLeft(line, "fold along ")
			steps = append(steps, step)
		}
	}
	fmt.Println("Num dots, xMax, yMax: ", len(dots), xMax, yMax)
	canvas := make([][]bool, yMax+1)
	for i, _ := range canvas {
		canvas[i] = make([]bool, xMax+1)
	}
	for _, d := range dots {
		canvas[d.y][d.x] = true
	}
	return canvas, steps
}
