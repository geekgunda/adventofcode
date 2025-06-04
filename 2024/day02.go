//go:build ignore

package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidReport(report []int) bool {
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
			return false
		}
	}
	return true
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
	var safeReportCount, dampenedReportCount int
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
		if isValidReport(parsedReport) {
			//log.Printf("Report is safe: %v", report)
			safeReportCount++
		} else {
			var isValid bool
			for i := 0; i < len(parsedReport); i++ {
				var rdelta []int
				rdelta = append(rdelta, parsedReport[:i]...)
				rdelta = append(rdelta, parsedReport[i+1:]...)
				if isValidReport(rdelta) {
					dampenedReportCount++
					isValid = true
					break
				}
			}
			if !isValid {
				log.Printf("Invalid report. Report: %v", parsedReport)
			}
		}
	}
	log.Printf("Result for part 1: %v", safeReportCount)
	log.Printf("Result for part 2: %v", safeReportCount+dampenedReportCount)
}
