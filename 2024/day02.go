package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidReport(report []int) (bool, int) {
	var isIncreasing bool
	for i := 1; i < len(report); i++ {
		if i == 1 {
			if report[i] > report[i-1] {
				isIncreasing = true
			}
		}
		diff := report[i] - report[i-1]
		if (isIncreasing && diff >= 1 && diff <= 3) ||
			(!isIncreasing && diff <= -1 && diff >= -3) {
			continue
		} else {
			//log.Printf("Invalid report: Index: %v | Report: %v", i, report)
			return false, i
		}
	}
	return true, 0
}

func main() {
	// example input
	//input := "ex2.txt"
	// exercise input
	input := "2.txt"
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var safeReportCount, dampenedRecordCount int
	var parsedReport []int
	for report := range strings.Lines(string(bytes)) {
		levels := strings.Split(report, " ")
		parsedReport = nil
		for i, level := range levels {
			parsed, err := strconv.Atoi(strings.Trim(level, "\n"))
			if err != nil {
				log.Fatalf("Error parsing index: %v | Error: %v", i, err)
			}
			parsedReport = append(parsedReport, parsed)
		}
		if valid, index := isValidReport(parsedReport); valid {
			//log.Printf("Report is safe: %v", report)
			safeReportCount++
		} else {
			// Skip the previous level
			var rdelta []int
			rdelta = append(rdelta, parsedReport[:index-1]...)
			rdelta = append(rdelta, parsedReport[index:]...)
			if v, _ := isValidReport(rdelta); v {
				//log.Printf("Valid dampened report Index: %v | Report: %v", index, parsedReport)
				dampenedRecordCount++
				continue
			}
			// Skip the current level
			rdelta = nil
			if index+1 < len(parsedReport) {
				rdelta = append(rdelta, parsedReport[:index]...)
				rdelta = append(rdelta, parsedReport[index+1:]...)
			} else {
				// Index is the last element
				rdelta = append(rdelta, parsedReport[:index]...)
			}
			if v, _ := isValidReport(rdelta); v {
				//log.Printf("Valid dampened report Index: %v | Report: %v", index, parsedReport)
				dampenedRecordCount++
				continue
			}
			log.Printf("Invalid report. Index: %v | Report: %v", index, parsedReport)
		}
	}
	log.Printf("Result for part 1: %v", safeReportCount)
	log.Printf("Result for part 2: %v", safeReportCount+dampenedRecordCount)
}
