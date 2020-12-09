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
	res, pos := FindInvalidNumber(input, 25)
	log.Printf("(Part 1) First invalid number in the sequence: %d", res)
	res = FindEncryptionWeakness(input, pos)
	log.Printf("(Part 2) Encryption weakness: %d", res)
}

func FindInvalidNumber(numbers []int64, preambleLen int64) (int64, int64) {
	var i, res, size int64
	preamble := make(map[int64]bool)
	for i = 0; i < preambleLen; i++ {
		preamble[numbers[i]] = true
	}
	size = int64(len(numbers))
	for i = preambleLen; i < size; i++ {
		preamble[numbers[i-1]] = true
		//log.Printf("Preamble: %#v", preamble)
		isValid := FindAdditives(numbers[i], preamble)
		if !isValid {
			res = numbers[i]
			break
		}
		//log.Printf("Next num: %d | Found pieces: %d and %d", numbers[i], n1, n2)
		delete(preamble, numbers[i-preambleLen])
	}
	return res, i
}

func FindAdditives(num int64, pieces map[int64]bool) bool {
	for n, _ := range pieces {
		if _, ok := pieces[num-n]; ok && n != num-n {
			return true
		}
	}
	return false
}

func FindEncryptionWeakness(numbers []int64, invalidPos int64) (res int64) {
	target := numbers[invalidPos]
	p1, p2, isMatch := FindContiguousAdditives(numbers, target)
	if isMatch {
		n1, n2 := FindSmallestAndLargest(numbers[p1 : p2+1])
		res = n1 + n2
		log.Printf("Found match at pos: %d to %d | values: %d and %d | sum: %d", p1, p2, n1, n2, res)
	}
	return
}

func FindContiguousAdditives(numbers []int64, target int64) (p1, p2 int64, isMatch bool) {
	count := int64(len(numbers))
	for i := int64(1); i < count; i++ {
		for j := int64(0); j < count-i; j++ {
			s := sum(numbers[j : j+i+1])
			log.Printf("Trying to find contiguous additives: len: %d | start: %d | end: %d | sum: %d | target: %d", i, j, j+i, s, target)
			if s == target {
				p1 = j
				p2 = j + i
				isMatch = true
				break
			}
		}
		if isMatch {
			break
		}
	}
	return
}

func sum(numbers []int64) (res int64) {
	for _, n := range numbers {
		res += n
	}
	return
}

func FindSmallestAndLargest(nums []int64) (int64, int64) {
	var s, l int64
	s = nums[0]
	l = nums[0]
	for _, n := range nums {
		if n < s {
			s = n
		}
		if n > l {
			l = n
		}
	}
	return s, l
}
