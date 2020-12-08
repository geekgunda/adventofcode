package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(\w+) (\W\w+)`)
}

func main() {
	inputFile := "../input/d8.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	insts := ParseInstructions(lines)
	res, _ := FindLoop(insts)
	log.Printf("(Part 1) Accumulator value before loop: %d", res)
	res = FixLoop(insts)
	log.Printf("(Part 2) Accumulator value after fixing loop: %d", res)
}

type Instruction struct {
	Cmd  string
	Step int
}

func ParseInstructions(commands []string) []Instruction {
	var insts []Instruction
	for i := 0; i < len(commands); i++ {
		m := re.FindStringSubmatch(commands[i])
		if len(m) != 3 {
			log.Printf("Invalid pattern: %s", commands[i])
			continue
		}
		step, err := strconv.Atoi(m[2])
		if err != nil {
			log.Printf("Error parsing step: %v", err)
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
			log.Printf("Loop detected at pos: %d", i)
			isLoop = true
			break
		}
		log.Printf("Processed command: %s | step: %v | pos: %d | accumulator: %d", insts[i].Cmd, insts[i].Step, i, res)
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
			log.Printf("Tried replacing %d inst jmp -> nop ", i)
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
			log.Printf("Tried replacing %d inst nop -> jmp ", i)
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
