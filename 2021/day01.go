package main

func day1() error {
	numbers, err := readFileAsInts64()
	if err != nil {
		return err
	}
	res := singleMeasurement(numbers)
	logResult(1, 1, "Result is", res)
	res = threeMeasurement(numbers)
	logResult(1, 2, "Result is", res)
	return nil
}

func singleMeasurement(input []int64) int64 {
	res := int64(0)
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			res++
		}
	}
	return res
}

func threeMeasurement(input []int64) int64 {
	var currSum, prevSum, res int64
	if len(input) < 3 {
		return res
	}
	prevSum = input[0] + input[1] + input[2]
	for i := 1; i < len(input)-2; i++ {
		currSum = input[i] + input[i+1] + input[i+2]
		if currSum > prevSum {
			res++
		}
		prevSum = currSum
	}
	return res
}
