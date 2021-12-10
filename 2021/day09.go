package main

import (
	"fmt"
	"sort"
)

func day9() error {
	input, err := readFileAs2DIntArray()
	if err != nil {
		return err
	}
	sum := sumLowPoints(input)
	logResult(9, 1, "Sum of risk level of low points: ", sum)
	product := findLargestBasins(input)
	logResult(9, 2, "Product of top 3 basins: ", product)
	return nil
}

func findLargestBasins(heights [][]int) int {
	var sizes sort.IntSlice
	for i, row := range heights {
		for j, h := range row {
			// Ignore points with height of 9
			if h == 9 {
				continue
			}
			lookup := make(map[string]bool)
			size := findBasinSize(heights, lookup, i, j)
			sizes = append(sizes, size)
		}
	}
	sort.Sort(sizes)
	l := len(sizes)
	fmt.Println("Top 4 basins size: ", sizes[l-5:l])
	return sizes[l-1] * sizes[l-2] * sizes[l-3]
}

func findBasinSize(heights [][]int, lookup map[string]bool, x, y int) int {
	key := fmt.Sprintf("%d_%d", x, y)
	var res int
	height := heights[x][y]
	// Ignore points with height 9
	if height == 9 {
		return 0
	}
	// Only consider this point, if we haven't seen it already for current root coordinates
	if _, ok := lookup[key]; !ok {
		res++
	}
	// Add it for next lookup
	lookup[key] = true
	//fmt.Printf("%d", height)
	x1, x2, y1, y2 := x-1, x+1, y-1, y+1
	if x1 >= 0 && heights[x1][y] > height {
		res += findBasinSize(heights, lookup, x1, y)
	}
	if x2 < len(heights) && heights[x2][y] > height {
		res += findBasinSize(heights, lookup, x2, y)
	}
	if y1 >= 0 && heights[x][y1] > height {
		res += findBasinSize(heights, lookup, x, y1)
	}
	if y2 < len(heights[x]) && heights[x][y2] > height {
		res += findBasinSize(heights, lookup, x, y2)
	}
	return res
}

func sumLowPoints(heights [][]int) int {
	var sum int
	for i, row := range heights {
		for j, num := range row {
			isLow := true
			// check if it's a low point
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					// ignore the point and diagonals
					if x == y || x+y == 0 {
						continue
					}
					a := i + x
					b := j + y
					// check bounds
					if a < 0 || a >= len(heights) || b < 0 || b >= len(row) {
						continue
					}
					if heights[a][b] <= num {
						isLow = false
						break
					}
				}
				if !isLow {
					break
				}
			}
			if isLow {
				sum += num + 1
			}
		}
	}
	return sum
}
