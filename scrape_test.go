package main

import (
	"reflect"
	"testing"
)

func TestSign(t *testing.T) {
	type test struct {
		input  string
		output int
	}
	tests := []test{
		{input: "Aries\n", output: 1},
		{input: "Virgo\n", output: 6},
		{input: "Libra\n", output: 7},
		{input: "Leo\n", output: 5},
		{input: "Boob\n", output: 0},
	}
	for _, tc := range tests {
		got := getMonthOfBirth(tc.input)
		if !reflect.DeepEqual(tc.output, got) {
			t.Fatalf("expected %v, got: %v", tc.output, got)
		}
	}
}
