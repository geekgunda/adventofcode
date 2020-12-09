package main

import (
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
	log.Println("Total passes found: ", len(lines))
	highestSeatID := GetHighestSeatID(lines)
	log.Printf("(Part 1) Highest Seat ID: %d", highestSeatID)
	missingSeatID := FindMissingSeat(lines)
	log.Printf("(Part 2) Missing Seat ID: %d", missingSeatID)
}

func FindMissingSeat(passes []string) int {
	seats := make(map[int]string)
	// make a map of all possible seats
	var found, high, low int
	low = 1024
	log.Println("Total boarding passes: ", len(passes))
	for _, pass := range passes {
		seatID := GetSeatIDFromBoardingPass(pass)
		seats[seatID] = pass
		found++
		if seatID < low {
			low = seatID
		}
		if seatID > high {
			high = seatID
		}
	}
	log.Printf("Count of remaining seats: %d | High: %d | Low: %d | Found: %d", high-low-found+1, high, low, found)
	for i := low + 1; i < high; i++ {
		if _, ok := seats[i]; !ok { // if this seat is missing
			if _, ok = seats[i-1]; ok { // if seat -1 is present
				if _, ok = seats[i+1]; ok { // if seat +1 is present
					return i
				}
			}
		}
	}
	return 0
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
	var r rune
	for _, r = range pass[:7] {
		n /= 2
		if r == 'F' {
			max -= n
		} else if r == 'B' {
			min += n
		} else {
			log.Printf("Invalid char: %s | Pass: %s", string(r), pass)
		}
	}
	row = min
	n = 8
	min = 0
	max = 7
	for _, r := range pass[7:] {
		n /= 2
		if r == 'L' {
			max -= n
		} else if r == 'R' {
			min += n
		} else {
			log.Printf("Invalid char: %s | Pass: %s", string(r), pass)
		}
	}
	col = min
	return (row * 8) + col
}
