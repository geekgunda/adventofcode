package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile := "../input/d1.txt"
	result := int64(2020)
	numbers, err := ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	product, err := FindMatchingPair(numbers, result)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Product(2) is: %v", product)
	if product, err = FindMatchingTriplet(numbers, result); err != nil {
		log.Fatalf("Error finding triplet: %v", err)
	}
	log.Printf("Product(3) is: %v", product)

}

func ReadFile(inputFile string) ([]int64, error) {
	var f *os.File
	var err error
	f, err = os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var numbers []int64
	var line []byte
	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		var num int64
		if num, err = strconv.ParseInt(string(line), 10, 64); err == nil {
			numbers = append(numbers, num)
		} else {
			return nil, err
		}
	}
	return numbers, nil
}

func FindMatchingPair(numbers []int64, result int64) (int64, error) {
	lookup := make(map[int64]bool)
	for _, n := range numbers {
		if _, ok := lookup[result-n]; ok {
			return n * (result - n), nil
		}
		lookup[n] = true
	}
	return int64(0), fmt.Errorf("Not found")
}

func FindMatchingTriplet(numbers []int64, result int64) (int64, error) {
	lookup := make(map[int64]bool)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if _, ok := lookup[result-numbers[i]-numbers[j]]; ok {
				return (numbers[i] * numbers[j] * (result - numbers[i] - numbers[j])), nil
			}
			lookup[numbers[i]] = true
			lookup[numbers[j]] = true
		}
	}
	return int64(0), fmt.Errorf("Not found")
}
