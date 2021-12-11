package main

import "fmt"

func day11() error {
	input, err := readFileAs2DIntArray()
	if err != nil {
		return err
	}
	flashes, allFlashStep := countFlashes(input, 100)
	logResult(11, 1, "Total flashes: ", flashes)
	logResult(11, 2, "All flash step: ", allFlashStep)
	return nil
}

func countFlashes(energyLevels [][]int, steps int) (int, int) {
	var flashes int
	fmt.Println()
	var elementCount int
	for _, row := range energyLevels {
		elementCount += len(row)
	}
	allFlash := false
	step := 0
	for allFlash == false {
		// first increment everyone's energy by 1
		for i, row := range energyLevels {
			for j, energyLevel := range row {
				energyLevels[i][j] = energyLevel + 1
			}
		}
		// keep track of positions that flashed in this step
		flashMap := make(map[string]bool)
		// If energy is greater than 9 and neighbours haven't been flashed, do it!
		stableState := false
		for !stableState {
			stableState = true
			for i, row := range energyLevels {
				for j, energyLevel := range row {
					if energyLevel > 9 {
						// Avoid duplicate flashes
						key := fmt.Sprintf("%d_%d", i, j)
						if _, ok := flashMap[key]; !ok {
							stableState = false
							flashMap[key] = true
							energyLevels, flashMap = flashNeighbours(energyLevels, flashMap, i, j)
						}
					}
				}
			}
		}
		// reset energy level of all flashed octopuses and print current state
		if (step+1)%10 == 0 {
			fmt.Println(step + 1)
		}
		for i, row := range energyLevels {
			for j, e := range row {
				if e > 9 {
					if step < steps {
						flashes++
					}
					energyLevels[i][j] = 0
					if (step+1)%10 == 0 {
						fmt.Print("*")
					}
				} else if (step+1)%10 == 0 {
					fmt.Printf("%d", e)
				}
			}
			if (step+1)%10 == 0 {
				fmt.Println()
			}
		}
		if (step+1)%10 == 0 {
			fmt.Println()
		}
		if len(flashMap) == elementCount {
			allFlash = true
		}
		step++
	}
	return flashes, step
}

func flashNeighbours(energyLevels [][]int, flashMap map[string]bool, i, j int) ([][]int, map[string]bool) {
	// iterate through all neighbours and increment their energy
	for x := i - 1; x <= i+1; x++ {
		// check bounds
		if x < 0 || x >= len(energyLevels) {
			continue
		}
		for y := j - 1; y <= j+1; y++ {
			// check bounds
			if y < 0 || y >= len(energyLevels[x]) {
				continue
			}
			// ignore the point itself
			if x == i && y == j {
				continue
			}
			energyLevels[x][y] += 1
			// If energy is greater than 9 and neighbours haven't been flashed, do it!
			/*
				if energyLevels[x][y] > 9 {
					key := fmt.Sprintf("%d_%d", x, y)
					if _, ok := flashMap[key]; !ok {
						flashMap[key] = true
						energyLevels, flashMap = flashNeighbours(energyLevels, flashMap, x, y)
					}
				}*/
		}
	}
	return energyLevels, flashMap
}
