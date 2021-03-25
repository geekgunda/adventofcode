package main

func day11() error {
	seats, err := readFileAsStrings()
	if err != nil {
		return err
	}
	seatCount := FindStableLayoutSeatCount(seats, 4, 1)
	logResult(11, 1, "Seat count", seatCount)
	seatCount = FindStableLayoutSeatCount(seats, 5, 0)
	logResult(11, 2, "Seat count", seatCount)
	return nil
}

func FindStableLayoutSeatCount(seats []string, threshold, nestLevel int) int {
	var i int
	for {
		i++
		nextIteration, isChanged := ApplySeatingRules(seats, threshold, nestLevel)
		//fmt.Printf("After iter: %d | Result: \n", i)
		/*for _, r := range nextIteration {
			fmt.Printf("%s\n", r)
		}*/
		if !isChanged {
			break
		}
		seats = nextIteration
	}
	var count int
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func ApplySeatingRules(seats []string, threshold, nestLevel int) ([]string, bool) {
	var isChanged = false
	result := make([]string, len(seats))
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			// Ignore floor
			if seats[i][j] == '.' {
				result[i] += "."
				continue
			}
			var occupancy int
			if nestLevel == 0 {
				occupancy = getFirstSeatOccupancy(seats, i, j)
			} else {
				occupancy = getOccupancy(seats, i, j, nestLevel)
			}
			//fmt.Printf("Occupancy for pos %d,%d is: %d\n", i, j, occupancy)
			if seats[i][j] == 'L' && occupancy == 0 {
				isChanged = true
				result[i] += "#"
			} else if seats[i][j] == '#' && occupancy >= threshold {
				isChanged = true
				result[i] += "L"
			} else {
				result[i] += string(seats[i][j])
			}
		}
	}
	return result, isChanged
}

func getOccupancy(seats []string, row, col, nestLevel int) int {
	var count int
	for i := row - nestLevel; i <= row+nestLevel; i++ {
		if i < 0 || i >= len(seats) {
			//fmt.Println("Ignoring row: ", i)
			continue
		}
		for j := col - nestLevel; j <= col+nestLevel; j++ {
			if j < 0 || j >= len(seats[row]) {
				//fmt.Printf("Ignoring row: %d | col: %d\n", i, j)
				continue
			}
			if i == row && j == col {
				continue
			}
			if seats[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func getFirstSeatOccupancy(seats []string, row, col int) int {
	var count int
	// horizontal left
	for x := row - 1; x >= 0; x-- {
		if seats[x][col] != '.' {
			if seats[x][col] == '#' {
				count++
			}
			break
		}
	}
	// horizontal right
	for x := row + 1; x < len(seats); x++ {
		if seats[x][col] != '.' {
			if seats[x][col] == '#' {
				count++
			}
			break
		}
	}
	// vertically up
	for y := col - 1; y >= 0; y-- {
		if seats[row][y] != '.' {
			if seats[row][y] == '#' {
				count++
			}
			break
		}
	}
	// vertically down
	for y := col + 1; y < len(seats[row]); y++ {
		if seats[row][y] != '.' {
			if seats[row][y] == '#' {
				count++
			}
			break
		}
	}
	// diagonal top left
	for i := -1; row+i >= 0; i-- {
		if col+i < 0 {
			break
		}
		if seats[row+i][col+i] != '.' {
			if seats[row+i][col+i] == '#' {
				count++
			}
			break
		}
	}
	// diagonal bottom right
	for i := 1; row+i < len(seats); i++ {
		if col+i >= len(seats[row]) {
			break
		}
		if seats[row+i][col+i] != '.' {
			if seats[row+i][col+i] == '#' {
				count++
			}
			break
		}
	}
	// diagonal top right
	for i := -1; row+i >= 0; i-- {
		if col-i >= len(seats[row]) {
			break
		}
		if seats[row+i][col-i] != '.' {
			if seats[row+i][col-i] == '#' {
				count++
			}
			break
		}
	}
	// diagonal bottom left
	for i := 1; row+i < len(seats); i++ {
		if col-i < 0 {
			break
		}
		if seats[row+i][col-i] != '.' {
			if seats[row+i][col-i] == '#' {
				count++
			}
			break
		}
	}
	return count
}
