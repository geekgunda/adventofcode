package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputFile := "../input/d9.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	input := make([]int64, len(lines))
	for i, l := range lines {
		tmp, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			log.Fatalf("Error parsing number: %s | Err: %v", l, err)
		}
		input[i] = tmp
	}
	res := FindInvalidNumber(input, 25)
	log.Printf("(Part 1) First invalid number in the sequence: %d", res)
}

func FindInvalidNumber(numbers []int64, preambleLen int64) int64 {
	var i, res, size int64
	preamble := make(map[int64]bool)
	for i = 0; i < preambleLen; i++ {
		preamble[numbers[i]] = true
	}
	size = int64(len(numbers))
	for i = preambleLen; i < size; i++ {
		preamble[numbers[i-1]] = true
		log.Printf("Preamble: %#v", preamble)
		n1, n2, isValid := FindAdditives(numbers[i], preamble)
		if !isValid {
			res = numbers[i]
			break
		}
		log.Printf("Next num: %d | Found pieces: %d and %d", numbers[i], n1, n2)
		delete(preamble, numbers[i-preambleLen])
	}
	return res
}

func FindAdditives(num int64, pieces map[int64]bool) (n1, n2 int64, isValid bool) {
	for n, _ := range pieces {
		if _, ok := pieces[num-n]; ok && n != num-n {
			n1 = n
			n2 = num - n
			isValid = true
			break
		}
	}
	return
}
