package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "slices"
)

// do this all the time in game dev for tilemaps
// if we have a 9 x 9 matrix we can easily calculate/track the offsets
// [-1,  1, 1]
// [-1,  0, 1]
// [-1, -1, 1]
//
// e.g.
//
// . . @ @ @ .
// . . . @ @ @
// . . @ @ @ .
//
// [ [0,0] 1,0 2,0 3,0 4,0 5,0
//   0,1 1,1 2,1 3,1 4,1 5,1
//   0,2 1,2 2,2 3,2 4,2 5,2
// ]
//
// to check for [3, 1]  we apply the matrix to the array and check [-1, -1]:[2, 0], [-1, 0]: [2, 1] etc.
// if we have a @ we increment the rollCounter if the rollCounter > 3 we stop and skip
// for that node if we are ok we add the roll as solution roll

var tilemap = [][]int{
	{-1, 1}, {0, 1}, {1, 1},
	{-1, 0}, {0, 0}, {1, 0},
	{-1, -1}, {0, -1}, {1, -1},
}

type Coord struct {
	X int
	Y int
}

type Checker struct {
	CheckUpLeft   Coord
	CheckLeft     Coord
	CheckBtmLeft  Coord
	CheckUpRight  Coord
	CheckRight    Coord
	CheckBtmRight Coord
	CheckUp       Coord
	CheckBtm      Coord
	AllXs         []int
	AllYs         []int
}

func InitChecker(checkX int, checkY int) Checker {
	AllXs := []int{}
	AllYs := []int{}

	checkUpLeft := Coord{X: checkX + tilemap[0][0], Y: checkY + tilemap[0][1]}
	AllYs = append(AllYs, checkUpLeft.Y)
	AllXs = append(AllXs, checkUpLeft.X)

	checkLeft := Coord{X: checkX + tilemap[3][0], Y: checkY + tilemap[3][1]}
	AllYs = append(AllYs, checkLeft.Y)
	// AllXs = append(AllXs, checkLeft.X)

	checkBtmLeft := Coord{X: checkX + tilemap[6][0], Y: checkY + tilemap[6][1]}
	AllYs = append(AllYs, checkBtmLeft.Y)
	// AllXs = append(AllXs, checkBtmLeft.X)

	checkUpRight := Coord{X: checkX + tilemap[2][0], Y: checkY + tilemap[2][1]}
	// AllYs = append(AllYs, checkUpRight.Y)
	AllXs = append(AllXs, checkUpRight.X)

	checkRight := Coord{X: checkX + tilemap[5][0], Y: checkY + tilemap[5][1]}
	// AllYs = append(AllYs, checkRight.Y)
	// AllXs = append(AllXs, checkRight.X)

	checkBtmRight := Coord{X: checkX + tilemap[8][0], Y: checkY + tilemap[8][1]}
	// AllYs = append(AllYs, checkBtmRight.Y)
	// AllXs = append(AllXs, checkBtmRight.X)

	checkUp := Coord{X: checkX + tilemap[1][0], Y: checkY + tilemap[1][1]}
	// AllYs = append(AllYs, checkUp.Y)
	AllXs = append(AllXs, checkUp.X)

	checkBtm := Coord{X: checkX + tilemap[7][0], Y: checkY + tilemap[7][1]}
	// AllYs = append(AllYs, checkBtm.Y)
	// AllXs = append(AllXs, checkBtm.X)

	return Checker{
		CheckUpLeft:   checkUpLeft,
		CheckLeft:     checkLeft,
		CheckBtmLeft:  checkBtmLeft,
		CheckUpRight:  checkUpRight,
		CheckRight:    checkRight,
		CheckBtmRight: checkBtmRight,
		CheckUp:       checkUp,
		CheckBtm:      checkBtm,
		AllXs:         AllXs,
		AllYs:         AllYs,
	}
}

func lookForRolls(input [][]string, checkX int, checkY int) bool {
	checker := InitChecker(checkX, checkY)
	// fmt.Println(checker)
	rolls := 0
	for _, y := range checker.AllYs {
		if y >= 0 && len(input) > y {
			// fmt.Println(y)
			row := input[y]
			// fmt.Println(row)
			for _, x := range checker.AllXs {
				if x >= 0 && len(row) > x {
					// fmt.Println(x)
					val := row[x]
					// fmt.Println(val)
					if val == "@" {
						rolls += 1
					}
				}
			}
		}
	}
	rolls = rolls - 1
	// fmt.Println(rolls)
	// if rolls < 4 {
	// 	fmt.Printf("coords {x: %d, y: %d}\n", checkX, checkY)
	// }
	// fmt.Println(rolls)
	return rolls < 4
}

func lookForRollsPt2(input [][]string, checkX int, checkY int) bool {
	checker := InitChecker(checkX, checkY)
	// fmt.Println(checker)
	rolls := 0
	for _, y := range checker.AllYs {
		if y >= 0 && len(input) > y {
			// fmt.Println(y)
			row := input[y]
			// fmt.Println(row)
			for _, x := range checker.AllXs {
				if x >= 0 && len(row) > x {
					// fmt.Println(x)
					val := row[x]
					// fmt.Println(val)
					if val == "@" {
						rolls += 1
					}
				}
			}
		}
	}
	rolls = rolls - 1
	// fmt.Println(rolls)
	// if rolls < 4 {
	// 	fmt.Printf("coords {x: %d, y: %d}\n", checkX, checkY)
	// }
	// fmt.Println(rolls)
	return rolls < 4
}

func coordsToCheck(input [][]string) []Coord {
	coordsToCheck := []Coord{}
	for x, line := range input {
		for y, val := range line {
			if val == "@" {
				coordsToCheck = append(coordsToCheck, Coord{X: x, Y: y})
			}
		}
	}
	// fmt.Println(coordsToCheck)
	return coordsToCheck
}

func solve1(input [][]string) int {
	accessible := 0
	for y, line := range input {
		for x, val := range line {
			if val == "@" {
				if lookForRolls(input, x, y) {
					accessible += 1
				}
			}
		}
	}
	return accessible
}

func solve2(input [][]string, totalRolls int) int {
	accessible := 0
	for y, line := range input {
		for x, val := range line {
			if val == "@" {
				if lookForRollsPt2(input, x, y) {
					input[y][x] = "."
					accessible += 1
				}
			}
		}
	}
	totalRolls = totalRolls + accessible
	if accessible == 0 {
		return totalRolls
	}
	return solve2(input, totalRolls)
}

func parseFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	arrays := [][]string{}
	for _, line := range lines {
		arrays = append(arrays, strings.Split(line, ""))
	}
	return arrays
}

func main() {
	input := parseFile("./input.txt")
	accessible := solve2(input, 0)
	fmt.Println(accessible)
}
