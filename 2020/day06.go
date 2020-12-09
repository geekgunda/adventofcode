package main

func day6() error {
	lines, err := readFileAsStrings()
	if err != nil {
		return err
	}
	anyCount, everyCount := GetCustomsAnswerCount(lines)
	logResult(6, 1, "Total Customs answer count (anyone answered)", anyCount)
	logResult(6, 2, "Total Customs answer count (everyone answered)", everyCount)
	return nil
}

func GetCustomsAnswerCount(input []string) (int, int) {
	var anyCount, everyCount, groupCount int
	groupAns := make(map[rune]int)
	for _, ans := range input {
		if len(ans) == 0 {
			anyCount += len(groupAns)
			for _, ct := range groupAns {
				if ct == groupCount {
					everyCount++
				}
			}
			//log.Printf("Group Answers: %#v | Count: %d", groupAns, count)
			groupAns = map[rune]int{}
			groupCount = 0
			continue
		}
		for _, r := range ans {
			if _, ok := groupAns[r]; ok {
				groupAns[r]++
			} else {
				groupAns[r] = 1
			}
		}
		groupCount++
	}
	if len(groupAns) > 0 {
		anyCount += len(groupAns)
		for _, ct := range groupAns {
			if ct == groupCount {
				everyCount++
			}
		}
	}
	return anyCount, everyCount
}
