package tasks

import (
	"fmt"
	"testing"
)

func TestIsCorrectBracketsSeq(t *testing.T) {
	inputs := map[string]bool{
		"{(()}":  false,
		"()":     true,
		"{[()]}": true,
		"()[]{}": true,
		"(]":     false,
		"(":      false,
		"]":      false,
	}
	for input, expected := range inputs {
		fmt.Println(input)
		fmt.Printf("%v, expected: %v\n", IsCorrectBracketsSeq(input), expected)
	}
}
