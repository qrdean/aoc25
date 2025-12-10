package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		bounds := []Bounds{
			{low: 3, high: 5},
			{low: 10, high: 14},
			{low: 16, high: 20},
			{low: 12, high: 18},
		}

		testNumbs := []int{
			1, 5, 8, 11, 17, 32,
		}
		got := solve1(testNumbs, bounds)
		want := 3
		if got != want {
			t.Errorf("got: %d != want: %d", got, want)
		}
	})

	t.Run("solve pt2", func(t *testing.T) {
		bounds := []Bounds{
			{low: 3, high: 5},
			{low: 10, high: 14},
			{low: 16, high: 20},
			{low: 12, high: 18},
		}
		got := solve2(bounds)
		want := 14
		if got != want {
			t.Errorf("got: %d != want: %d", got, want)
		}
	})

	t.Run("parse input", func(t *testing.T) {
		input, bounds :=	parseInputs("./test.txt", "./fresh_ranges_test.txt") 
		wantInput := []int{
			1, 5, 8, 11, 17, 32,
		}
		wantBounds := []Bounds{
			{low: 3, high: 5},
			{low: 10, high: 14},
			{low: 16, high: 20},
			{low: 12, high: 18},
		}
		if !reflect.DeepEqual(input, wantInput) {
			t.Errorf("input: %v, != want: %v", input, wantInput)
		}

		if !reflect.DeepEqual(bounds, wantBounds) {
			t.Errorf("bounds: %v, != want: %v", bounds, wantBounds)
		}
	})

}

func BenchmarkPt1(b *testing.B) {
	input, bounds := parseInputs("./input.txt", "./fresh_ranges_input.txt")
  b.ResetTimer()
	solve1(input, bounds)
}

func BenchmarkPt2(b *testing.B) {
	_, bounds := parseInputs("./input.txt", "./fresh_ranges_input.txt")
	b.ResetTimer()
	solve2(bounds)
}
