package main

import (
	"fmt"
	"strconv"
)

func day3() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	gammaRate, epsilonRate := countForBits(input)
	var res int64
	if res, err = getResult(gammaRate, epsilonRate); err != nil {
		return err
	}
	logResult(3, 1, "Power consumption is: ", res)
	o2Rate, co2Rate := calculateRatings(input)
	if res, err = getResult(o2Rate, co2Rate); err != nil {
		return err
	}
	logResult(3, 2, "Life support rating is: ", res)
	return nil
}

func getResult(x, y string) (int64, error) {
	var xInt, yInt int64
	var err error
	if xInt, err = strconv.ParseInt(x, 2, 64); err != nil {
		return xInt, err
	}
	if yInt, err = strconv.ParseInt(y, 2, 64); err != nil {
		return yInt, err
	}
	return xInt * yInt, nil
}

func countForBits(input []string) (string, string) {
	zeroCount := make([]int, len(input[0]))
	oneCount := make([]int, len(input[0]))
	for _, line := range input {
		for i, b := range line {
			switch b {
			case '0':
				zeroCount[i]++
			case '1':
				oneCount[i]++
			default:
				fmt.Errorf("Invalid input: %v", b)
			}
		}
	}
	var mostCommon, leastCommon string
	for i := 0; i < len(zeroCount); i++ {
		if zeroCount[i] > oneCount[i] {
			mostCommon += "0"
			leastCommon += "1"
		} else {
			mostCommon += "1"
			leastCommon += "0"
		}
	}
	return mostCommon, leastCommon
}

func calculateRatings(input []string) (string, string) {
	var o2Rating, co2Rating string
	// O2 Rating
	o2input := input
	for i := 0; i < len(input[0]); i++ {
		o2input = applyBitCriteria(o2input, i, true)
		if len(o2input) == 1 {
			o2Rating = o2input[0]
			break
		}
	}
	// CO2 Rating
	co2input := input
	for i := 0; i < len(input[0]); i++ {
		co2input = applyBitCriteria(co2input, i, false)
		if len(co2input) == 1 {
			co2Rating = co2input[0]
			break
		}
	}
	return o2Rating, co2Rating
}

func applyBitCriteria(input []string, pos int, mostCommon bool) []string {
	var zeroRatings, oneRatings []string
	for _, line := range input {
		switch line[pos] {
		case '0':
			zeroRatings = append(zeroRatings, line)
		case '1':
			oneRatings = append(oneRatings, line)
		}
	}
	if len(zeroRatings) == len(oneRatings) {
		if mostCommon {
			return oneRatings
		} else {
			return zeroRatings
		}
	} else if len(zeroRatings) > len(oneRatings) {
		if mostCommon {
			return zeroRatings
		} else {
			return oneRatings
		}
	} else {
		if mostCommon {
			return oneRatings
		} else {
			return zeroRatings
		}
	}
}
