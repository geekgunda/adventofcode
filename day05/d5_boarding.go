package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inputFile := "../input/d5.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	highestSeatID := GetHighestSeatID(lines)
	log.Printf("(Part 1) Highest Seat ID: %d", highestSeatID)
}

func GetHighestSeatID(passes []string) int {
	var result int
	for _, pass := range passes {
		newSeatID := GetSeatIDFromBoardingPass(pass)
		if newSeatID > result {
			result = newSeatID
		}
	}
	return result
}
func GetSeatIDFromBoardingPass(pass string) int {
	var row, col, min, max, n int
	n = 128
	max = n - 1
	min = 0
	for _, r := range pass[:6] {
		n /= 2
		if r == 'F' {
			max -= n
		} else {
			min += n
		}
		//fmt.Printf("Row | r: %s | min: %d | max: %d\n", string(r), min, max)
	}
	row = min
	n = 8
	min = 0
	max = 7
	for _, r := range pass[7:] {
		n /= 2
		if r == 'L' {
			max -= n
		} else {
			min += n
		}
		//fmt.Printf("Col | r: %s | min: %d | max: %d\n", string(r), min, max)
	}
	col = min
	fmt.Printf("Pass: %s | row: %d | col: %d\n", pass, row, col)
	return (row * 8) + col
}
