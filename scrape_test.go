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

// I would like to make a second test table for the horoscope
// but each horoscope would change through the day so I wont be able to test that function
// sorry

// Cancer
func BenchmarkCancer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scrapeHoro(4)
	}
}

// Capricorn
func BenchmarkCapricorn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scrapeHoro(10)
	}
}

// Taurus
func BenchmarkTaurus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scrapeHoro(2)
	}
}

// Libra
func BenchmarkLibra(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scrapeHoro(2)
	}
}
