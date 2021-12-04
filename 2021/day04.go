package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day4() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	numbers := parseNumbers(input[0])
	boards := parseBoards(input[1:])
	score := runBingo(numbers, boards, true)
	logResult(4, 1, "Score for first winning board is: ", score)
	score = runBingo(numbers, boards, false)
	logResult(4, 2, "Score for last winning board is: ", score)
	return nil
}

func runBingo(bingoNumbers []int, boards [][][]int, returnFirstWin bool) int {
	var score int
	boardStatus := make([]bool, len(boards))
	var winCount int
	for _, bingoNum := range bingoNumbers {
		for boardNum, board := range boards {
			for rowNum, row := range board {
				for colNum, num := range row {
					if num == bingoNum {
						// mark it as seen
						boards[boardNum][rowNum][colNum] = -1
						// check if board won
						if checkIfBoardWon(board) {
							//fmt.Println("Winning board: ", board)
							sum := scoreBoard(board)
							score = sum * bingoNum
							winCount++
							if winCount == 1 && returnFirstWin {
								return score
							}
							boardStatus[boardNum] = true
							// check if all boards won
							var boardWinCount int
							for _, won := range boardStatus {
								if won {
									boardWinCount++
								}
							}
							if boardWinCount == len(boards) {
								return score
							}
						}
					}
				}
			}
		}
	}
	return score
}

func checkIfBoardWon(board [][]int) bool {
	// check for row win
	var sum int
	for _, row := range board {
		for _, num := range row {
			sum += num
		}
		if sum == -5 {
			//fmt.Println("Horizontal match")
			return true
		}
		sum = 0
	}
	sum = 0
	// check for column win
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board); j++ {
			sum += board[j][i]
		}
		if sum == -5 {
			//fmt.Println("Vertical match")
			return true
		}
		sum = 0
	}
	return false
}

func scoreBoard(board [][]int) int {
	var sum int
	for _, row := range board {
		for _, num := range row {
			if num == -1 {
				continue
			}
			sum += num
		}
	}
	return sum
}

func parseNumbers(numbers string) []int {
	numArr := strings.Split(numbers, ",")
	var res []int
	for _, num := range numArr {
		n, e := strconv.Atoi(num)
		if e != nil {
			fmt.Errorf("Error parsing num: %v", e)
		}
		res = append(res, n)
	}
	return res
}

func parseBoards(boards []string) [][][]int {
	var parsedBoards [][][]int
	var board [][]int
	var boardRow []int
	var num string
	for _, l := range boards {
		if len(l) == 0 {
			// old board ends
			if len(board) > 0 && len(board[0]) > 0 {
				parsedBoards = append(parsedBoards, board)
				//fmt.Println("Board is: ", board)
			}
			// reset the current board
			board = [][]int{}
			continue
		}
		for _, c := range l {
			if c == ' ' {
				if len(num) > 0 {
					n, err := strconv.Atoi(num)
					if err != nil {
						fmt.Errorf("Error converting number: %v", err)
					}
					boardRow = append(boardRow, n)
					num = ""
				}
			} else {
				// append the digit into string placeholder
				num += string(c)
			}
		}
		//last number in each line
		if len(num) > 0 {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Errorf("Error converting number: %v", err)
			}
			boardRow = append(boardRow, n)
			num = ""
		}
		// add boardRow to board
		board = append(board, boardRow)
		boardRow = []int{}
	}
	return parsedBoards
}
