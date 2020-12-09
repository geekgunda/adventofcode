package main

import "fmt"

func day5() error {
	lines, err := readFileAsStrings()
	if err != nil {
		return err
	}
	highestSeatID := GetHighestSeatID(lines)
	logResult(5, 1, "Highest Seat ID", highestSeatID)
	missingSeatID := FindMissingSeat(lines)
	logResult(5, 2, "Missing Seat ID", missingSeatID)
	return nil
}

func FindMissingSeat(passes []string) int {
	seats := make(map[int]string)
	// make a map of all possible seats
	var found, high, low int
	low = 1024
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
	//fmt.Printf("Count of remaining seats: %d | High: %d | Low: %d | Found: %d\n", high-low-found+1, high, low, found)
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
			fmt.Printf("Invalid char: %s | Pass: %s\n", string(r), pass)
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
			fmt.Printf("Invalid char: %s | Pass: %s\n", string(r), pass)
		}
	}
	col = min
	return (row * 8) + col
}
