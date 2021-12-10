package main

import (
	"fmt"
	"sort"
)

func day10() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	score, cScore := parseAndScoreSyntax(input)
	logResult(10, 1, "Score for syntax error : ", score)
	logResult(10, 2, "Score for auto-correct : ", cScore)
	return nil
}

func parseAndScoreSyntax(input []string) (int, int) {
	var score int
	var correctionScore sort.IntSlice
	closingChunk := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	for i, line := range input {
		var stack []string
		var corrupted bool
		//fmt.Println("Line: ", line)
		for j, r := range line {
			//fmt.Println()
			//fmt.Printf("char: %s", string(r))
			if cl, ok := closingChunk[string(r)]; ok {
				stack = append(stack, cl)
				//fmt.Printf("\tadded to stack: %v", stack)
				continue
			}
			if len(stack) > 0 {
				//fmt.Println("Found char: ", string(r), " stack: ", stack)
				expected := stack[len(stack)-1]
				if expected == string(r) {
					stack = stack[:len(stack)-1]
					//fmt.Printf("\tpop from stack: %v", stack)
					continue
				} else {
					// ignore the corrupted line
					corrupted = true
					score += getScoreForError(string(r))
					fmt.Println("Corrupted at line ", i, " position ", j, " expected: ", expected, " score: ", score)
					break
				}
			}
			// we didn't find the closing chunk
		}
		// check for incomplete chunks
		if len(stack) > 0 && !corrupted {
			fmt.Println("Incomplete line: ", i, " stack: ", stack)
			// start with a score of 0 for each line
			var cScore int
			// completion needs to happen in reverse order, a.k.a pop from stack
			for i := len(stack) - 1; i >= 0; i-- {
				c := stack[i]
				cScore *= 5                        // multiply score by 5 for each completion
				cScore += getScoreForCompletion(c) // add score for completion of this chunk
				//fmt.Println("char: ", c, " score: ", cScore)
			}
			correctionScore = append(correctionScore, cScore)
		}
	}
	sort.Sort(correctionScore)      // sort scores
	mid := len(correctionScore) / 2 // find median
	fmt.Println("Correction scores: ", correctionScore)
	return score, correctionScore[mid]
}

func getScoreForError(key string) int {
	var res int
	switch key {
	case ")":
		res = 3
	case "]":
		res = 57
	case "}":
		res = 1197
	case ">":
		res = 25137
	default:
		fmt.Println("Invalid character: ", key)
	}
	return res
}

func getScoreForCompletion(key string) int {
	var res int
	switch key {
	case ")":
		res = 1
	case "]":
		res = 2
	case "}":
		res = 3
	case ">":
		res = 4
	default:
		fmt.Println("Invalid completion character: ", key)
	}
	return res
}
