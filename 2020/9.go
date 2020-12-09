package main

func day9() error {
	input, err := readFileAsInts64()
	if err != nil {
		return err
	}
	res, pos := FindInvalidNumber(input, 25)
	logResult(9, 1, "First invalid number in the sequence", res)
	res = FindEncryptionWeakness(input, pos)
	logResult(9, 2, "Encryption weakness", res)
	return nil
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
		isValid := FindAdditives(numbers[i], preamble)
		if !isValid {
			res = numbers[i]
			break
		}
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
		//fmt.Printf("Found match at pos: %d to %d | values: %d and %d | sum: %d\n", p1, p2, n1, n2, res)
	}
	return
}

func FindContiguousAdditives(numbers []int64, target int64) (p1, p2 int64, isMatch bool) {
	count := int64(len(numbers))
	for i := int64(1); i < count; i++ {
		for j := int64(0); j < count-i; j++ {
			s := sum(numbers[j : j+i+1])
			//fmt.Printf("Trying to find contiguous additives: len: %d | start: %d | end: %d | sum: %d | target: %d\n", i, j, j+i, s, target)
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
