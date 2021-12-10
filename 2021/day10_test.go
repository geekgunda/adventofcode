package main

import "testing"

func TestDay10SyntaxErrorScore(t *testing.T) {
	var input = []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
	score, cScore := parseAndScoreSyntax(input)
	if score != 26397 {
		t.Errorf("Invalid score: %d", score)
	}
	if cScore != 288957 {
		t.Errorf("Invalid correction score: %d", cScore)
	}
}
