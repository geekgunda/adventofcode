package main

import (
	"fmt"
	"strings"
)

func day12() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	paths, pathsV2 := countCavePaths(input)
	logResult(12, 1, "Count of paths: ", paths)
	logResult(12, 2, "Count of paths (with double traversal): ", pathsV2)
	return nil
}

type Cave struct {
	name                    string
	isSmall, isStart, isEnd bool
	connIter                int
	connections             []*Cave
}

func (c *Cave) addConnection(newConn *Cave) {
	// check if this connection already exists
	// if it does, we have nothing to do
	for _, conn := range c.connections {
		if conn.name == newConn.name {
			return
		}
	}
	// Don't add a path TO start cave
	if newConn.isStart {
		return
	}
	// Don't add a path FROM end cave
	if c.isEnd {
		return
	}
	// if not, add the connection
	c.connections = append(c.connections, newConn)
}

func NewCave(name string) *Cave {
	return &Cave{
		name:    name,
		isSmall: name == strings.ToLower(name),
		isStart: name == "start",
		isEnd:   name == "end",
	}
}

func countCavePaths(input []string) (int, int) {
	caveSystem := parseCaveSystem(input)
	fmt.Println("Number of caves: ", len(caveSystem))
	var paths int
	/**
	 * A map doesn't work for this use-case because maps are passed by reference in go
	 * Because of this property we carry forward the traversals of small caves, across
	 * different paths, which messes up the count
	 */
	//visited := make(map[string]int)
	visited := []string{}
	paths = countPaths(caveSystem["start"], nil, visited, 0)
	fmt.Println("Count with no double traversals: ", paths)
	pathsV2 := paths
	for _, c := range caveSystem {
		if c.isStart || c.isEnd {
			continue
		}
		if !c.isSmall {
			continue
		}
		visited = []string{}
		pathsV2 += (countPaths(caveSystem["start"], c, visited, 0) - paths)
		fmt.Println("Count with traversing cave twice: ", c.name, pathsV2)
	}
	return paths, pathsV2
}

func parseCaveSystem(input []string) map[string]*Cave {
	lookup := make(map[string]*Cave)
	for _, line := range input {
		// parse the line into individual caves
		ends := strings.Split(line, "-")
		if len(ends) != 2 {
			fmt.Println("Error parsing cave connection: ", line)
			continue
		}
		c1, c2 := ends[0], ends[1]
		// create or lookup this particular cave
		var cave1, cave2 *Cave
		if val, ok := lookup[c1]; ok {
			cave1 = val
		} else {
			cave1 = NewCave(c1)
			lookup[c1] = cave1
		}
		if val, ok := lookup[c2]; ok {
			cave2 = val
		} else {
			cave2 = NewCave(c2)
			lookup[c2] = cave2
		}
		// add connections between these caves
		cave1.addConnection(cave2)
		cave2.addConnection(cave1)
	}
	return lookup
}

func countPaths(cave, doubleTraversalCave *Cave, visited []string, level int) int {
	level++
	var count int
	//fmt.Println(strings.Repeat("\t", level-1), cave.name)
	if cave.isEnd {
		//fmt.Println()
		count++
		return count
	}
	if cave.isSmall && !cave.isEnd && !cave.isStart {
		if isCaveVisited(cave, doubleTraversalCave, visited) == false {
			visited = append(visited, cave.name)
		}
	}
	for _, conn := range cave.connections {
		if conn.isSmall && !conn.isEnd && !cave.isStart {
			if isCaveVisited(conn, doubleTraversalCave, visited) {
				continue
			}
		}
		count += countPaths(conn, doubleTraversalCave, visited, level)
	}
	return count
}

func isCaveVisited(cave, doubleTraversal *Cave, visited []string) bool {
	var seenCount int
	for _, name := range visited {
		// if the requested cave is found in the list
		if cave.name == name {
			// and double traversal is not allowed, cave visits are done
			if doubleTraversal == nil {
				return true
			}
			// if this cave is not the one we traverse twice, cave visits are done
			if cave.name != doubleTraversal.name {
				return true
			}
			// otherwise, allow traversing this cave twice
			seenCount++
			if seenCount == 2 {
				return true
			}
			continue
		}
	}
	return false
}
