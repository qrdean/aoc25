package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Bounds struct {
	low  int
	high int
}

func checkRange(bounds []Bounds, num int) bool {
	for _, bound := range bounds {
		if bound.low <= num && bound.high >= num {
			return true
		}
	}
	return false
}

func solve1(input []int, bounds []Bounds) int {
	freshIngredient := 0
	for _, num := range input {
		if checkRange(bounds, num) {
			freshIngredient += 1
		}
	}
	return freshIngredient
}

func sortMin(a Bounds, b Bounds) int {
	if a.low < b.low {
		return -1
	}
	if a.low > b.low {
		return 1
	}
	return 0
}

func accumulateHighest(highest int, bound Bounds) (int, int) {
	if bound.high < highest {
		return 0, highest
	}
	total := bound.high - max(highest, bound.low) + 1
	highest = bound.high + 1
	return total, highest
}

func solve2(bounds []Bounds) int {
	totalIngredients := 0
	highest := 0
	slices.SortFunc(bounds, sortMin)
	for _, bound := range bounds {
		var total int
		total, highest = accumulateHighest(highest, bound)
		totalIngredients += total
	}
	return totalIngredients
}

func parseInputs(ingredientPath string, boundsPath string) ([]int, []Bounds) {
	file, err := os.Open(ingredientPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	arrays := []int{}
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		arrays = append(arrays, num)
	}
	file2, err := os.Open(boundsPath)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	scanner = bufio.NewScanner(file2)
	lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	bounds := []Bounds{}
	for _, line := range lines {
		boundsStr := strings.Split(line, "-")
		low, _ := strconv.Atoi(boundsStr[0])
		high, _ := strconv.Atoi(boundsStr[1])
		bound := Bounds{low: low, high: high}
		bounds = append(bounds, bound)
	}
	return arrays, bounds
}

func main() {
	_, bounds := parseInputs("./input.txt", "./fresh_ranges_input.txt")
	res := solve2(bounds)
	fmt.Println(res)
}
