//go:build ignore

package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := "7.txt"
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading input file [%v]: %v", input, err)
	}
	eqs := parseInput(bytes)
	allowedOperations = []CalibrationOperation{Addition, Multiplication}
	log.Printf("Part 1 | Calibration result: %v", validateEquations(eqs))
	allowedOperations = []CalibrationOperation{Addition, Multiplication, Concatenation}
	log.Printf("Part 2 | Calibration result: %v", validateEquations(eqs))
}

func validateEquations(eqs []CalibrationEquation) int {
	var res int
	for _, eq := range eqs {
		if eq.validate(eq.numbers[0], 1) == true {
			res += eq.testVal
		}
	}
	return res
}

type CalibrationOperation int

const (
	Addition CalibrationOperation = iota + 1
	Multiplication
	Concatenation
)

var allowedOperations []CalibrationOperation

func (op CalibrationOperation) apply(a, b int) int {
	switch op {
	case Addition:
		return a + b
	case Multiplication:
		return a * b
	case Concatenation:
		if b == 0 {
			return a * 10
		}
		l10 := math.Log10(float64(b))
		exp := math.Pow10(int(l10) + 1)
		return a*int(exp) + b
	default:
		log.Fatalf("Invalid operation: %v", op)
	}
	return 0
}

type CalibrationEquation struct {
	testVal int
	numbers []int
}

func (eq CalibrationEquation) validate(currVal, i int) bool {
	//log.Printf("Validate: currVal: [%v] | Index: [%v] | Target: [%v]", currVal, i, eq.testVal)
	if i == len(eq.numbers) {
		return currVal == eq.testVal
	}
	var nextVal int
	for _, op := range allowedOperations {
		nextVal = op.apply(currVal, eq.numbers[i])
		if eq.validate(nextVal, i+1) == true {
			return true
		}
	}
	return false
}

func parseInput(bytes []byte) []CalibrationEquation {
	var res []CalibrationEquation
	for line := range strings.Lines(string(bytes)) {
		eq := CalibrationEquation{}
		l := strings.Trim(line, "\n")
		parts := strings.Split(l, ":")
		if len(parts) != 2 {
			log.Fatalf("Couldn't split equation [%v] into 2 parts", l)
		}
		val, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Error parsing test value [%v]: %v", parts[0], err)
		}
		eq.testVal = val
		//nums := strings.Split(parts[1], " ")
		for num := range strings.SplitSeq(parts[1], " ") {
			n := strings.TrimSpace(num)
			if len(n) == 0 {
				continue
			}
			val, err = strconv.Atoi(n)
			if err != nil {
				log.Fatalf("Error parsing number [%v]: %v", n, err)
			}
			eq.numbers = append(eq.numbers, val)
		}
		res = append(res, eq)
	}
	return res
}
