package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// t.Run("test me", func(t *testing.T) {
	// 	testInput := parseFile("./test.txt")
	//
	// 	successCoords := []Coord{
	// 		{X: 2, Y: 0},
	// 		{X: 3, Y: 0},
	// 		{X: 5, Y: 0},
	// 		{X: 6, Y: 0},
	// 		{X: 0, Y: 1},
	// 		{X: 6, Y: 2},
	// 		{X: 0, Y: 4},
	// 		{X: 9, Y: 4},
	// 		{X: 0, Y: 7},
	// 		{X: 0, Y: 9},
	// 		{X: 2, Y: 9},
	// 		{X: 8, Y: 9},
	// 	}
	// 	failureCoords := []Coord{
	// 		{X: 7, Y: 0},
	// 		{X: 1, Y: 1},
	// 		{X: 2, Y: 1},
	// 		{X: 4, Y: 1},
	// 		{X: 6, Y: 1},
	// 		{X: 8, Y: 1},
	// 		{X: 9, Y: 1},
	// 		{X: 0, Y: 2},
	// 		{X: 1, Y: 2},
	// 		{X: 2, Y: 2},
	// 		{X: 3, Y: 2},
	// 		{X: 4, Y: 2},
	// 		{X: 8, Y: 2},
	// 		{X: 9, Y: 2},
	// 		{X: 0, Y: 3},
	// 		{X: 2, Y: 3},
	// 		{X: 3, Y: 3},
	// 		{X: 4, Y: 3},
	// 		{X: 5, Y: 3},
	// 		{X: 8, Y: 3},
	// 		{X: 1, Y: 4},
	// 		{X: 3, Y: 4},
	// 		{X: 4, Y: 4},
	// 		{X: 5, Y: 4},
	// 		{X: 6, Y: 4},
	// 		{X: 8, Y: 4},
	// 		{X: 1, Y: 5},
	// 		{X: 2, Y: 5},
	// 		{X: 3, Y: 5},
	// 		{X: 4, Y: 5},
	// 		{X: 5, Y: 5},
	// 		{X: 6, Y: 5},
	// 		{X: 7, Y: 5},
	// 		{X: 9, Y: 5},
	// 		{X: 1, Y: 6},
	// 		{X: 3, Y: 6},
	// 		{X: 5, Y: 6},
	// 		{X: 7, Y: 6},
	// 		{X: 8, Y: 6},
	// 		{X: 9, Y: 6},
	// 		{X: 2, Y: 7},
	// 		{X: 3, Y: 7},
	// 		{X: 4, Y: 7},
	// 		{X: 6, Y: 7},
	// 		{X: 7, Y: 7},
	// 		{X: 8, Y: 7},
	// 		{X: 9, Y: 7},
	// 		{X: 1, Y: 8},
	// 		{X: 2, Y: 8},
	// 		{X: 3, Y: 8},
	// 		{X: 4, Y: 8},
	// 		{X: 5, Y: 8},
	// 		{X: 6, Y: 8},
	// 		{X: 7, Y: 8},
	// 		{X: 8, Y: 8},
	// 		{X: 4, Y: 9},
	// 		{X: 5, Y: 9},
	// 		{X: 6, Y: 9},
	// 	}
	//
	// 	for _, coord := range successCoords {
	// 		if !lookForRolls(testInput, coord.X, coord.Y) {
	// 			t.Errorf("x:%d y:%d,should work", coord.X, coord.Y)
	// 		}
	// 	}
	// 	for _, coord := range failureCoords {
	// 		if lookForRolls(testInput, coord.X, coord.Y) {
	// 			t.Errorf("x:%dy:%d,should not work", coord.X, coord.Y)
	// 		}
	// 	}
	// })

	t.Run("test coords to check is correct", func(t *testing.T) {
		input := parseFile("./test.txt")
		got := coordsToCheck(input)
		want := []Coord{
			{X: 2, Y: 0},
			{X: 3, Y: 0},
			{X: 5, Y: 0},
			{X: 6, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: 1},
			{X: 6, Y: 2},
			{X: 0, Y: 4},
			{X: 9, Y: 4},
			{X: 0, Y: 7},
			{X: 0, Y: 9},
			{X: 2, Y: 9},
			{X: 8, Y: 9},
			{X: 7, Y: 0},
			{X: 1, Y: 1},
			{X: 2, Y: 1},
			{X: 4, Y: 1},
			{X: 6, Y: 1},
			{X: 8, Y: 1},
			{X: 9, Y: 1},
			{X: 0, Y: 2},
			{X: 1, Y: 2},
			{X: 2, Y: 2},
			{X: 3, Y: 2},
			{X: 4, Y: 2},
			{X: 8, Y: 2},
			{X: 9, Y: 2},
			{X: 0, Y: 3},
			{X: 2, Y: 3},
			{X: 3, Y: 3},
			{X: 4, Y: 3},
			{X: 5, Y: 3},
			{X: 8, Y: 3},
			{X: 1, Y: 4},
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 6, Y: 4},
			{X: 8, Y: 4},
			{X: 1, Y: 5},
			{X: 2, Y: 5},
			{X: 3, Y: 5},
			{X: 4, Y: 5},
			{X: 5, Y: 5},
			{X: 6, Y: 5},
			{X: 7, Y: 5},
			{X: 9, Y: 5},
			{X: 1, Y: 6},
			{X: 3, Y: 6},
			{X: 5, Y: 6},
			{X: 7, Y: 6},
			{X: 8, Y: 6},
			{X: 9, Y: 6},
			{X: 2, Y: 7},
			{X: 3, Y: 7},
			{X: 4, Y: 7},
			{X: 6, Y: 7},
			{X: 7, Y: 7},
			{X: 8, Y: 7},
			{X: 9, Y: 7},
			{X: 1, Y: 8},
			{X: 2, Y: 8},
			{X: 3, Y: 8},
			{X: 4, Y: 8},
			{X: 5, Y: 8},
			{X: 6, Y: 8},
			{X: 7, Y: 8},
			{X: 8, Y: 8},
			{X: 4, Y: 9},
			{X: 5, Y: 9},
			{X: 6, Y: 9},
		}
		if reflect.DeepEqual(got, want) {
			t.Errorf("want: %d != got %d", want, got)
		}
	})

	t.Run("test input", func(t *testing.T) {
		input := parseFile("./test.txt")
		got := solve1(input)
		want := 13
		if want != got {
			t.Errorf("want: %d != got %d", want, got)
		}
	})

	t.Run("test input pt2", func(t *testing.T) {
		input := parseFile("./test.txt")
		got := solve2(input, 0)
		want := 43
		if want != got {
			t.Errorf("want: %d != got %d", want, got)
		}
	})

	t.Run("testsplit", func(t *testing.T) {
		testInput := "..@@.@@@@."
		want := []string{".", ".", "@", "@", ".", "@", "@", "@", "@", "."}
		got := strings.Split(testInput, "")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %v != got: %v", want, got)
		}
	})
}

func BenchmarkPt1(b *testing.B) {
	testInput := parseFile("./input.txt")
	b.ResetTimer()
	solve1(testInput)
}


func BenchmarkPt2(b *testing.B) {
	testInput := parseFile("./input.txt")
	b.ResetTimer()
	solve2(testInput, 0)
}
