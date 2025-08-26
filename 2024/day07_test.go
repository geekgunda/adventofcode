package main

import (
	"fmt"
	"testing"
)

// Run as: go test -v day07.go day07_test.go

func TestDay7Part1(t *testing.T) {
	input := []struct {
		eq  CalibrationEquation
		res bool
	}{
		{eq: CalibrationEquation{testVal: 190, numbers: []int{10, 19}}, res: true},
		{eq: CalibrationEquation{testVal: 3267, numbers: []int{81, 40, 27}}, res: true},
		{eq: CalibrationEquation{testVal: 83, numbers: []int{17, 5}}, res: false},
		{eq: CalibrationEquation{testVal: 156, numbers: []int{15, 6}}, res: false},
		{eq: CalibrationEquation{testVal: 7290, numbers: []int{6, 8, 6, 15}}, res: false},
		{eq: CalibrationEquation{testVal: 161011, numbers: []int{16, 10, 13}}, res: false},
		{eq: CalibrationEquation{testVal: 192, numbers: []int{17, 8, 14}}, res: false},
		{eq: CalibrationEquation{testVal: 21037, numbers: []int{9, 7, 18, 13}}, res: false},
		{eq: CalibrationEquation{testVal: 292, numbers: []int{11, 6, 16, 20}}, res: true},
	}
	allowedOperations = []CalibrationOperation{Addition, Multiplication}
	for i, tt := range input {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			res := tt.eq.validate(tt.eq.numbers[0], 1)
			if res != tt.res {
				t.Errorf("Test failed: Expected: %v | Actual: %v", tt.res, res)
			}
		})
	}
}

func TestDay7Part2(t *testing.T) {
	input := []struct {
		eq  CalibrationEquation
		res bool
	}{
		{eq: CalibrationEquation{testVal: 190, numbers: []int{10, 19}}, res: true},
		{eq: CalibrationEquation{testVal: 3267, numbers: []int{81, 40, 27}}, res: true},
		{eq: CalibrationEquation{testVal: 83, numbers: []int{17, 5}}, res: false},
		{eq: CalibrationEquation{testVal: 156, numbers: []int{15, 6}}, res: true},
		{eq: CalibrationEquation{testVal: 7290, numbers: []int{6, 8, 6, 15}}, res: true},
		{eq: CalibrationEquation{testVal: 161011, numbers: []int{16, 10, 13}}, res: false},
		{eq: CalibrationEquation{testVal: 192, numbers: []int{17, 8, 14}}, res: true},
		{eq: CalibrationEquation{testVal: 21037, numbers: []int{9, 7, 18, 13}}, res: false},
		{eq: CalibrationEquation{testVal: 292, numbers: []int{11, 6, 16, 20}}, res: true},
	}
	allowedOperations = []CalibrationOperation{Addition, Multiplication, Concatenation}
	for i, tt := range input {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			res := tt.eq.validate(tt.eq.numbers[0], 1)
			if res != tt.res {
				t.Errorf("Test failed: Expected: %v | Actual: %v", tt.res, res)
			}
		})
	}
}
