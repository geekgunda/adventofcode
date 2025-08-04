//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction int

const (
	DirUp Direction = iota + 1
	DirRight
	DirDown
	DirLeft
)

// changeDirection gives the next direction to traverse when an obstruction is found
func (d Direction) changeDirection() Direction {
	d2 := Direction(d + 1)
	if d2 > DirLeft {
		d2 = DirUp
	}
	return d2
}

// getNextStep gives the increment needed to x,y coordinates, to move to the next step in current direction
func (d Direction) getNextStep() (int, int) {
	switch d {
	case DirUp:
		return -1, 0
	case DirRight:
		return 0, 1
	case DirDown:
		return 1, 0
	case DirLeft:
		return 0, -1
	}
	log.Fatalf("Unknown direction: %v", d)
	return 0, 0
}

// dirText converts text notation into Direction type
var dirText = map[string]Direction{
	"^": DirUp,
	">": DirRight,
	"v": DirDown,
	"<": DirLeft,
}

type pos struct {
	x, y int
	dir  Direction
}

func main() {
	input := "6.txt"
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading file [%v]: %v", input, err)
	}
	var inputMatrix []string
	// Used to track the initial position of the guard
	var xi, yi int
	yi = -1
	var dir Direction
	var start pos
	for line := range strings.Lines(string(bytes)) {
		l := strings.Trim(line, "\n")
		inputMatrix = append(inputMatrix, l)
		// Find the initial position of the guard
		if yi == -1 {
			for s, d := range dirText {
				if yi = strings.Index(l, s); yi != -1 {
					start = pos{x: xi, y: yi, dir: d}
					break
				}
			}
			if yi == -1 {
				xi++
			}
		}
	}
	log.Printf("Starting position: x: [%v] y: [%v] dir: [%v]", xi, yi, dir)
	visitCount := traverseMatrix(inputMatrix, start)
	log.Printf("Part 1 | Visit count: %d", visitCount)
	obsCount := countObstructions(inputMatrix, start)
	log.Printf("Part 2 | Obstruction count: %d", obsCount)
}

// traverses given matrix at the starting point specified by x,y in the dir Direction
func traverseMatrix(input []string, start pos) int {
	// Track the distinct positions visited by the guard
	visitMap := make(map[string]bool)
	// Track the starting position
	x := start.x
	y := start.y
	dir := start.dir
	visitMap[fmt.Sprintf("%d_%d", x, y)] = true
	for x < len(input) && y < len(input[x]) {
		//log.Printf("x:[%v] y:[%v] dir:[%v]", x, y, dir)
		i, j := dir.getNextStep()
		xi := x + i
		yi := y + j
		//log.Printf("x:[%v] y:[%v] dir:[%v] next step", xi, yi, dir)
		// Check if next step is outside the bounds of the matrix. If so, we're done!
		if xi < 0 || xi >= len(input) || yi < 0 || yi >= len(input[xi]) {
			break
		}
		// Check if next step in the matrix is an obstruction. If so, change direction
		if input[xi][yi] == '#' {
			dir = dir.changeDirection()
			//log.Printf("x:[%v] y:[%v] dir:[%v] change dir", xi, yi, dir)
		} else {
			// If all other checks pass, traverse to the next step and capture it as visited
			x = xi
			y = yi
			// We don't want to count same position twice, if it has already been traversed
			visitMap[fmt.Sprintf("%d_%d", x, y)] = true
		}
	}
	return len(visitMap)
}

// counts obstructions that end up leading to traversal of matrix in a loop
func countObstructions(input []string, start pos) int {
	obsMap := make(map[string]bool)
	// Track the starting position
	x := start.x
	y := start.y
	dir := start.dir
	for x < len(input) && y < len(input[x]) {
		//log.Printf("x:[%v] y:[%v] dir:[%v]", x, y, dir)
		i, j := dir.getNextStep()
		xi := x + i
		yi := y + j
		//log.Printf("x:[%v] y:[%v] dir:[%v] next step", xi, yi, dir)
		// Check if next step is outside the bounds of the matrix. If so, we're done!
		if xi < 0 || xi >= len(input) || yi < 0 || yi >= len(input[xi]) {
			break
		}
		// Check if next step in the matrix is an obstruction. If so, change direction
		if input[xi][yi] == '#' {
			dir = dir.changeDirection()
			//log.Printf("x:[%v] y:[%v] dir:[%v] change dir", xi, yi, dir)
		} else {
			// Check if we end up in a loop if we put an obstruction at the next step
			//log.Printf("x:[%v] y:[%v] dir:[%v] obstacle at: x:[%v] y:[%v]", x, y, dir, xi, yi)
			if isLoop(input, start, pos{x: xi, y: yi}) {
				//log.Printf("x:[%v] y:[%v] dir:[%v] obstacle at: x:[%v] y:[%v] Success", x, y, dir, xi, yi)
				obsMap[fmt.Sprintf("%d_%d", xi, yi)] = true
			}
			// If all other checks pass, traverse to the next step and capture it as visited
			x = xi
			y = yi
		}
	}
	return len(obsMap)
}

// checks if putting an obstruction at a given position leads to a loop
func isLoop(input []string, start, obs pos) bool {
	x := start.x
	y := start.y
	dir := start.dir
	// Keep track of the visited positions
	visitMap := make(map[string]bool)
	for x < len(input) && y < len(input[x]) {
		i, j := dir.getNextStep()
		xi := x + i
		yi := y + j
		//log.Printf("x:[%v] y:[%v] dir:[%v] next step", xi, yi, dir)
		// Check if next step is outside the bounds of the matrix. If so, we're done!
		if xi < 0 || xi >= len(input) || yi < 0 || yi >= len(input[xi]) {
			break
		}
		// Check if next step in the matrix is an obstruction. If so, change direction
		if input[xi][yi] == '#' || (xi == obs.x && yi == obs.y) {
			dir = dir.changeDirection()
			//log.Printf("x:[%v] y:[%v] dir:[%v] change dir", xi, yi, dir)
		} else {
			// If all other checks pass, traverse to the next step and capture it as visited
			x = xi
			y = yi
			// Check if we've visited this position AND were heading in the same direction. If we have, it's a loop!
			visitKey := fmt.Sprintf("%d_%d_%v", x, y, dir)
			if _, ok := visitMap[visitKey]; ok {
				return true
			}
			visitMap[visitKey] = true
		}
	}
	return false
}
