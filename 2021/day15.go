package main

import (
	"container/heap"
	"math"
)

// This uses dijkstras algorithm to find shortest path
// Ref: https://medium.com/@verdi/understanding-dijkstras-algorithm-e1ce49c249f
func day15() error {
	input, err := readFileAs2DIntArray()
	if err != nil {
		return err
	}
	risk := findLowestRiskPath(input)
	logResult(15, 1, "Lowest risk: ", risk)
	fullRisk := findLowestRiskForFullCave(input)
	logResult(15, 2, "Lowest risk for full map: ", fullRisk)
	return nil
}

type CaveNode struct {
	x, y    int       // coordinates of this node
	risk    int       // risk value of this node
	cumRisk int       // cumulative risk value from start to this node
	prev    *CaveNode // points to previous node for shortest path
}

// Create a type to hold priority queue for this type
type Cavern []*CaveNode

// Implement heap.Interface to get a priority queue
func (c Cavern) Len() int           { return len(c) }                      // length of this queue
func (c Cavern) Less(i, j int) bool { return c[i].cumRisk < c[j].cumRisk } // determines the priority and which element to pop first
func (c Cavern) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }            // swap element at i and j
// Add a new element in queue
// DO NOT USE directly. Use heap.Push instead
func (c *Cavern) Push(x interface{}) {
	n := x.(*CaveNode)
	*c = append(*c, n)
}

// Pop the last element from queue
// Always pop LAST element. Less() will determine whether it's smallest or biggest
// DO NOT USE directly. Use heap.Pop instead
func (c *Cavern) Pop() interface{} {
	old := *c
	n := old[len(old)-1]
	old[len(old)-1] = nil
	*c = old[:len(old)-1]
	return n
}

func findLowestRiskForFullCave(input [][]int) int {
	fullCave := createFullCaveMap(input)
	return findLowestRiskPath(fullCave)
}

// Copy over a 2D matrix 'multiplier' times to both right and below
// Also increment the value by the copy number and reset >9 values to 1
func createFullCaveMap(input [][]int) [][]int {
	multiplier := 5
	fullMap := make([][]int, len(input)*multiplier)
	for i := 0; i < len(input)*multiplier; i++ {
		x := i
		var xinc int
		if x >= len(input) {
			x = i % len(input)    // relative position of this row relative to original matrix
			xinc = i / len(input) // copy number of this row
		}
		fullMap[i] = make([]int, len(input[x])*multiplier)
		for j := 0; j < len(input[x])*multiplier; j++ {
			y := j
			var yinc int
			if y >= len(input[x]) {
				y = j % len(input[x])    // relative position of column relative to original matrix
				yinc = j / len(input[x]) // copy number of this column
			}
			//fmt.Println(i, j, x, y)
			val := input[x][y]
			if xinc > 0 || yinc > 0 {
				val = val + xinc + yinc
				// If value is greater than 9 it resets to 1 (and NOT 0!)
				if val > 9 {
					val = val%10 + 1
				}
				//fmt.Println("Changed: ", input[x][y], " to ", val, " at ", i, j, xinc, yinc)
			}
			fullMap[i][j] = val
			//fmt.Print(val)
		}
		//fmt.Println()
	}
	return fullMap
}

func findLowestRiskPath(input [][]int) int {
	// 2-D array to hold node pointers for each 'point' in input
	allNodes := make([][]*CaveNode, len(input))
	// references to first and last nodes
	var start, end *CaveNode
	// First create node for each point in input
	for i, row := range input {
		allNodes[i] = make([]*CaveNode, len(row))
		for j, r := range row {
			node := &CaveNode{
				x:       i,
				y:       j,
				risk:    r,
				cumRisk: math.MaxInt64,
			}
			if i == 0 && j == 0 {
				start = node
			}
			if i == len(input)-1 && j == len(row)-1 {
				end = node
			}
			allNodes[i][j] = node
		}
	}
	// reset distance of first node as it's the starting point
	start.cumRisk = 0
	// This will store our priority queue
	var cavern Cavern
	// We start from first element
	heap.Push(&cavern, start)
	// This will initialize the heap invariants
	heap.Init(&cavern)
	for cavern.Len() > 0 {
		// because of priority queue implementation, we'll automatically get
		// the 'point' with shortest path from previous 'point'
		node := heap.Pop(&cavern).(*CaveNode)
		// If we are at end, we are done!
		if node == end {
			//fmt.Println("Popped end node")
			break
		}
		x, y := node.x, node.y
		//fmt.Println("Node: ", x, y)
		// We can only travel horizontally or vertically from this point
		// So edges for this node is it's neighbours, excluding diagonals
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				// ignore the point itself
				if i == 0 && j == 0 {
					continue
				}
				// ignore diagonal neighbours
				if i == j || i+j == 0 {
					continue
				}
				// Check bounds
				if x+i < 0 || x+i >= len(input) || y+j < 0 || y+j >= len(input[0]) {
					continue
				}
				edge := allNodes[x+i][y+j]
				riskFromCurNode := node.cumRisk + edge.risk
				// if we found a better path TO this neighbour, update it
				if edge.cumRisk > riskFromCurNode {
					edge.cumRisk = riskFromCurNode
					edge.prev = node
					// If we changed the priority of an element (a.k.a. it's cumulative risk), add it back in queue
					heap.Push(&cavern, edge)
				}
			}
		}
	}
	// Print the coordinates of the path
	/*
		fmt.Println("Path is: ")
		n := end
		for n != start {
			fmt.Printf("%d,%d->", n.x, n.y)
			n = n.prev
		}
		fmt.Println()
	*/
	return end.cumRisk
}
