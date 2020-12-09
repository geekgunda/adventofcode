package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day8() error {
	lines, err := readFileAsStrings()
	if err != nil {
		return err
	}
	insts := ParseInstructions(lines)
	res, _ := FindLoop(insts)
	logResult(8, 1, "Accumulator value before loop", res)
	res = FixLoop(insts)
	logResult(8, 2, "Accumulator value after fixing loop", res)
	return nil
}

type Instruction struct {
	Cmd  string
	Step int
}

func ParseInstructions(commands []string) []Instruction {
	var insts []Instruction
	re := regexp.MustCompile(`(\w+) (\W\w+)`)
	for i := 0; i < len(commands); i++ {
		m := re.FindStringSubmatch(commands[i])
		if len(m) != 3 {
			fmt.Printf("Invalid pattern: %s\n", commands[i])
			continue
		}
		step, err := strconv.Atoi(m[2])
		if err != nil {
			fmt.Printf("Error parsing step: %v\n", err)
			continue
		}
		insts = append(insts, Instruction{
			Cmd:  m[1],
			Step: step,
		})
	}
	return insts
}

func FindLoop(insts []Instruction) (int, bool) {
	var res int
	var isLoop bool
	lookup := make(map[int]bool)
	for i := 0; i < len(insts); {
		if _, ok := lookup[i]; ok {
			//fmt.Printf("Loop detected at pos: %d\n", i)
			isLoop = true
			break
		}
		//fmt.Printf("Processed command: %s | step: %v | pos: %d | accumulator: %d\n", insts[i].Cmd, insts[i].Step, i, res)
		lookup[i] = true
		switch insts[i].Cmd {
		case "acc":
			res += insts[i].Step
			i++
		case "jmp":
			i += insts[i].Step
		case "nop":
			i++
		}
	}
	return res, isLoop
}

func FixLoop(insts []Instruction) int {
	var res int
	for i := 0; i < len(insts); i++ {
		if insts[i].Cmd == "acc" {
			continue
		}
		if insts[i].Cmd == "jmp" {
			//fmt.Printf("Tried replacing %d inst jmp -> nop \n", i)
			newInsts := make([]Instruction, len(insts))
			copy(newInsts, insts)
			newInsts[i].Cmd = "nop"
			ct, isLoop := FindLoop(newInsts)
			if !isLoop {
				res = ct
				break
			}
			continue
		}
		if insts[i].Cmd == "nop" {
			//fmt.Printf("Tried replacing %d inst nop -> jmp \n", i)
			newInsts := make([]Instruction, len(insts))
			copy(newInsts, insts)
			newInsts[i].Cmd = "jmp"
			ct, isLoop := FindLoop(newInsts)
			if !isLoop {
				res = ct
				break
			}
			continue
		}
	}
	return res
}
