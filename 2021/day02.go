package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	x, y := calcCoordinates(input, false)
	logResult(2, 1, "Simple product is: ", x*y)
	x, y = calcCoordinates(input, true)
	logResult(2, 2, "Aimed product is: ", x*y)
	return nil
}

func calcCoordinates(input []string, useAim bool) (x, y int) {
	var aim int
	for _, c := range input {
		arr := strings.Split(c, " ")
		if len(arr) != 2 {
			fmt.Errorf("Invalid line: %v", c)
		}
		val, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Errorf("Error parsing command: %v - %v", c, err)
		}
		if useAim {
			aimedCalc(arr[0], val, &x, &y, &aim)
		} else {
			simpleCalc(arr[0], val, &x, &y)
		}
	}
	return
}

func simpleCalc(c string, v int, x, y *int) {
	switch c {
	case "forward":
		*x += v
	case "down":
		*y += v
	case "up":
		*y -= v
	default:
		fmt.Errorf("Invalid command: %v", c)
	}
}

func aimedCalc(c string, v int, x, y, a *int) {
	switch c {
	case "forward":
		*x += v
		*y += (*a) * v
	case "down":
		*a += v
	case "up":
		*a -= v
	default:
		fmt.Errorf("Invalid command: %v", c)
	}
}
