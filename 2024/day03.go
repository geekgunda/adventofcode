//go:build ignore

package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := "3.txt"
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var result, r2 int64
	enabled := true
	for line := range strings.Lines(string(bytes)) {
		l := strings.Trim(line, "\n")
		// Part 1
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re.FindAllStringSubmatch(l, -1)
		//log.Printf("Matches: %#v", matches)
		for _, sm := range matches {
			var n1, n2 int
			var err error
			// Ignore the first match that's the whole match
			if n1, err = strconv.Atoi(sm[1]); err != nil {
				log.Fatalf("Error parsing number: %v | Err: %v", sm[1], err)
			}
			if n2, err = strconv.Atoi(sm[2]); err != nil {
				log.Fatalf("Error parsing number: %v | Err: %v", sm[2], err)
			}
			//log.Printf("Numbers: %v * %v", n1, n2)
			result += int64(n1 * n2)
		}
		// Part 2
		re2 := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\((\d+),(\d+)\))`)
		m2 := re2.FindAllStringSubmatch(l, -1)
		//log.Printf("Matches: %#v", m2)
		// Instructions are enabled by default
		for i, sm := range m2 {
			//log.Printf("Index: %v | Submatch: %v", i, sm)
			switch {
			case len(sm[1]) > 0:
				//log.Printf("Enabled")
				enabled = true
				break
			case len(sm[2]) > 0:
				//log.Printf("Disabled")
				enabled = false
				break
			case len(sm[3]) > 0:
				if !enabled {
					break
				}
				var n1, n2 int
				var err error
				// Ignore the first match that's the whole match
				if n1, err = strconv.Atoi(sm[4]); err != nil {
					log.Fatalf("Error parsing number: %v | Err: %v", sm[4], err)
				}
				if n2, err = strconv.Atoi(sm[5]); err != nil {
					log.Fatalf("Error parsing number: %v | Err: %v", sm[5], err)
				}
				//log.Printf("Numbers: %v * %v", n1, n2)
				r2 += int64(n1 * n2)
			default:
				log.Fatalf("Index :%v | Default case triggered!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", i)
			}
		}
	}
	log.Printf("Result of part 1: %v", result)
	log.Printf("Result of part 2: %v", r2)
}
