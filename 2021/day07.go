package main

import (
	"fmt"
	"time"
)

func day7() error {
	t := time.Now()
	input, err := readFileAsNumbersCsv()
	if err != nil {
		return err
	}
	var minCost int
	// part 1
	minCost = findMinCost(input, true)
	logResult(7, 1, "Minimum fuel with contant rate: ", minCost)
	// part 2
	minCost = findMinCost(input, false)
	logResult(7, 2, "Minimum fuel with linear rate: ", minCost)
	fmt.Println("Time taken: ", time.Since(t))
	return nil
}

func findMinCost(positions []int, constantRate bool) int {
	var minPos, maxPos int
	minPos = positions[0]
	// Calculate min max bounds
	for _, pos := range positions {
		if minPos > pos {
			minPos = pos
		}
		if maxPos < pos {
			maxPos = pos
		}
	}
	//fmt.Println("Input size: ", len(positions), " min position: ", minPos, " max position: ", maxPos)
	// Calculate cost of moving to all possible locations and find the minimum cost
	var cost, minCost int
	for i := minPos; i < maxPos; i++ {
		cost = calculateCost(positions, i, constantRate)
		if minCost == 0 {
			minCost = cost
			//fmt.Println("Setting minCost first time to: ", minCost)
		}
		if cost < minCost {
			//fmt.Println("Found new minCost at position: ", i, " and cost: ", cost)
			minCost = cost
		}
	}
	return minCost
}

// Given a target position and individual positions of each submarine
// calculate total cost to move all submarines to target position
func calculateCost(positions []int, targetPosition int, constantRate bool) int {
	var totalCost, cost int
	for _, pos := range positions {
		if targetPosition > pos {
			if constantRate {
				cost = targetPosition - pos
			} else {
				cost = calculateLinearCost(pos, targetPosition)
			}
		} else {
			if constantRate {
				cost = pos - targetPosition
			} else {
				cost = calculateLinearCost(targetPosition, pos)
			}
		}
		totalCost += cost
	}
	return totalCost
}

// Calculate linearly increasing cost of moving from a start to end position
func calculateLinearCost(start, end int) int {
	var res, distance int
	for i := start; i < end; i++ {
		distance++
		res += distance
	}
	return res
}
