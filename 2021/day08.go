package main

import (
	"fmt"
	"sort"
	"strings"
)

func day8() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	displayItems := parseDisplayItems(input)
	var res int
	res = countUniqueDigits(displayItems)
	logResult(8, 1, "Count of unique digits: ", res)
	res = sumAllItems(displayItems)
	logResult(8, 2, "Sum of all displayed digits: ", res)
	return nil
}

type sortedString []rune

func (s sortedString) Len() int           { return len(s) }
func (s sortedString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortedString) Less(i, j int) bool { return s[i] < s[j] }

type DisplayItem struct {
	signals []string
	digits  []string
}

func parseDisplayItems(input []string) []DisplayItem {
	var res []DisplayItem
	for _, l := range input {
		parts := strings.Split(l, " | ")
		if len(parts) != 2 {
			fmt.Println("Error parsing input line: ", l)
		}
		signals := strings.Split(parts[0], " ")
		digits := strings.Split(parts[1], " ")
		if len(signals) != 10 {
			fmt.Println("Error parsing signals: ", signals, len(signals))
		}
		if len(digits) != 4 {
			fmt.Println("Error parsing digits: ", digits, len(digits))
		}
		var signalStrings, digitStrings []string
		// sort each signal alphabetically
		for _, s := range signals {
			ss := sortedString(s)
			sort.Sort(ss)
			signalStrings = append(signalStrings, string(ss))
		}
		// sort each digit alphabetically
		for _, s := range digits {
			ss := sortedString(s)
			sort.Sort(ss)
			digitStrings = append(digitStrings, string(ss))
		}
		item := DisplayItem{
			signals: signalStrings,
			digits:  digitStrings,
		}
		res = append(res, item)
	}
	return res
}

func countUniqueDigits(items []DisplayItem) int {
	var res int
	for _, item := range items {
		for _, digit := range item.digits {
			l := len(digit)
			if l == 2 {
				// 1
				res++
			} else if l == 4 {
				// 4
				res++
			} else if l == 3 {
				// 7
				res++
			} else if l == 7 {
				// 8
				res++
			}
		}
	}
	return res
}

func sumAllItems(items []DisplayItem) int {
	var res int
	for _, item := range items {
		num := decodeItem(item)
		res += num
	}
	return res
}

/**
 * Mapping of each digit, corresponding signal, and length of signal characters
 * (*) mark means the length is unique
 * 0 - abcefg (6)
 * 1 - cf (2)*
 * 2 - acdeg (5)
 * 3 - acdfg (5)
 * 4 - bcdf (4)*
 * 5 - abdfg (5)
 * 6 - abdefg (6)
 * 7 - acf (3)*
 * 8 - abcdefg (7)*
 * 9 - abcdfg (6)
 */
func decodeItem(item DisplayItem) int {
	// mapping of signal to it's corresponding digit
	// Adding 1 to each value, so that we can differentiate between missing signal
	// and the signal value for digit 0
	pendingSignals := make([]string, 6)
	signalMapping := make(map[string]int)
	var signal4, signal7 string
	for _, signal := range item.signals {
		l := len(signal)
		// Identify unique digits based on length first
		if l == 2 {
			// 1
			signalMapping[signal] = 1 + 1
		} else if l == 4 {
			// 4
			signal4 = signal
			signalMapping[signal] = 4 + 1
		} else if l == 3 {
			// 7
			signal7 = signal
			signalMapping[signal] = 7 + 1
		} else if l == 7 {
			// 8
			signalMapping[signal] = 8 + 1
		} else {
			pendingSignals = append(pendingSignals, signal)
		}
	}
	// We only have (3) 5-digit and (3) 6-digit sequences to decode
	for _, signal := range pendingSignals {
		strip4, strip7 := signal, signal
		// strip all signals from digit 4 from current signal
		for _, r := range signal4 {
			strip4 = strings.ReplaceAll(strip4, string(r), "")
		}
		// strip all signals from digit 7 from current signal
		for _, r := range signal7 {
			strip7 = strings.ReplaceAll(strip7, string(r), "")
		}
		l := len(signal)
		// identify 5-digit cases
		if l == 5 {
			/**
			 * 2 - acdeg (5)
			 * 3 - acdfg (5)
			 * 5 - abdfg (5)
			 * 4 - bcdf (4)*
			 * 7 - acf (3)*
			 */
			//   2     | 3     | 5
			//   acdeg | acdfg | abdfg
			// 4 aeg   | ag    | ag
			// 7 deg   | dg    | bdg
			if len(strip4) == 3 && len(strip7) == 3 {
				// 2
				signalMapping[signal] = 2 + 1
				continue
			}
			if len(strip4) == 2 && len(strip7) == 2 {
				// 3
				signalMapping[signal] = 3 + 1
				continue
			}
			if len(strip4) == 2 && len(strip7) == 3 {
				// 5
				signalMapping[signal] = 5 + 1
				continue
			}
		}
		if l == 6 {
			/**
			 * 0 - abcefg (6)
			 * 6 - abdefg (6)
			 * 9 - abcdfg (6)
			 * 4 - bcdf (4)*
			 * 7 - acf (3)*
			 */
			//   0      | 6      | 9
			//   abcefg | abdefg | abcdfg
			// 4 aeg    | aeg    | ag
			// 7 beg    | bdeg   | bdg
			if len(strip4) == 3 && len(strip7) == 3 {
				// 0
				signalMapping[signal] = 0 + 1
				continue
			}
			if len(strip4) == 3 && len(strip7) == 4 {
				// 6
				signalMapping[signal] = 6 + 1
				continue
			}
			if len(strip4) == 2 && len(strip7) == 3 {
				// 9
				signalMapping[signal] = 9 + 1
				continue
			}
		}
	}
	if len(signalMapping) != 10 {
		fmt.Println("Missing mappings: ", 10-len(signalMapping))
	}
	var res int
	factor := 1000
	for _, digit := range item.digits {
		val := signalMapping[digit]
		if val == 0 {
			fmt.Println("Missing interpretation for: ", digit)
		} else {
			res += (factor * (val - 1))
			factor /= 10
		}
	}
	return res
}
