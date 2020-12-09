package main

func day3() error {
	lines, err := readFileAsStrings()
	if err != nil {
		return err
	}
	count := GetTreeCount(lines, 1, 3)
	logResult(3, 1, "Trees encountered", count)
	slopes := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	product := GetMultipleTreeCount(lines, slopes)
	logResult(3, 2, "Product of trees encountered on slopes", product)
	return nil
}

func GetMultipleTreeCount(tobogganMap []string, slopes [][]int) int64 {
	product := int64(1)
	for _, row := range slopes {
		count := GetTreeCount(tobogganMap, row[0], row[1])
		//fmt.Printf("Count for slope: %v is %d", row, count)
		product *= int64(count)
	}
	return product
}

func GetTreeCount(tobogganMap []string, xInc, yInc int) int {
	var count, x, y int
	for x < len(tobogganMap) {
		if string(tobogganMap[x][y]) == "#" {
			count++
		}
		y = (y + yInc) % len(tobogganMap[x])
		x += xInc
	}
	return count
}
