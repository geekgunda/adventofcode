package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var TreeMap map[string]bool

func init() {
	TreeMap = make(map[string]bool)
}

func main() {
	inputFile := "../input/d3.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	count := GetTreeCount(lines, 1, 3)
	log.Printf("(Part 1) Trees encountered: %d", count)
	slopes := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	product := GetMultipleTreeCount(lines, slopes)
	log.Printf("(Part 2) Product of trees encountered on slopes: %d", product)
}

func GetMultipleTreeCount(tobogganMap []string, slopes [][]int) int64 {
	product := int64(1)
	for _, row := range slopes {
		count := GetTreeCount(tobogganMap, row[0], row[1])
		log.Printf("Count for slope: %v is %d", row, count)
		product *= int64(count)
	}
	return product
}

func GetTreeCount(tobogganMap []string, xInc, yInc int) int {
	var count, x, y int
	for x < len(tobogganMap) {
		if string(tobogganMap[x][y]) == "#" {
			count++
		}
		y = (y + yInc) % len(tobogganMap[x])
		x += xInc
	}
	return count
}
