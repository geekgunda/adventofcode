package main

import (
	"flag"
	"fmt"
)

var inputFile = flag.String("input", "", "input file relative path")
var problemNum = flag.Int("problem", 1, "problem number to run")

func main() {
	flag.Parse()
	if len(*inputFile) == 0 {
		*inputFile = fmt.Sprintf("%d.txt", *problemNum)
	}
	fmt.Printf("Solving problem [%d] with input file at: %s\n", *problemNum, *inputFile)
	if err := solveProblem(*problemNum); err != nil {
		fmt.Printf("Error while solving problem: %v\n", err)
	}
}

func solveProblem(num int) (err error) {
	switch num {
	case 1:
		err = day1()
	case 2:
		err = day2()
	case 3:
		err = day3()
	case 4:
		err = day4()
	case 5:
		err = day5()
	case 6:
		err = day6()
	case 7:
		err = day7()
	case 8:
		err = day8()
	case 9:
		err = day9()
	case 10:
		err = day10()
	case 11:
		err = day11()
	default:
		err = fmt.Errorf("This problem does not have a solution yet!")
	}
	return err
}
