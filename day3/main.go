package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BankNode struct {
	First  *BatteryNode
	Second *BatteryNode
}

func getCurrentJoltage(val1, val2 int) int {
	concatenate := fmt.Sprintf("%d%d", val1, val2)
	number, _ := strconv.Atoi(concatenate)
	return number
}

func getCurrentJoltageArr(vals [12]BatteryNode) int {
	concatenate := ""
	for _, val := range vals {
		concatenate += fmt.Sprintf("%d", val.Value)
	}
	number, _ := strconv.Atoi(concatenate)
	return number
}

type BatteryNode struct {
	Value    int
	Position int
}

func getBatteryValue(bank []int) int {
	if len(bank) == 0 {
		return 0
	}
	first := BatteryNode{Value: bank[len(bank)-2], Position: len(bank) - 2}
	second := BatteryNode{Value: bank[len(bank)-1], Position: len(bank) - 1}
	for i := len(bank) - 3; i >= 0; i-- {
		if first.Value <= bank[i] {
			temp := first.Value
			tempPos := first.Position
			first.Value = bank[i]
			first.Position = i

			if second.Value < temp {
				second.Value = temp
				second.Position = tempPos
			}
		}
	}
	return getCurrentJoltage(first.Value, second.Value)
}

func getBatteryValuePt2(bank []int) int {
	BatLimit := 12
	var batteryNodes [12]BatteryNode
	for i := range BatLimit {
		offset := BatLimit - i
		initPosition := len(bank) - offset
		batteryNodes[i] = BatteryNode{Value: bank[initPosition], Position: initPosition}
	}
	// fmt.Println(bank)
	for i := len(bank) - (BatLimit + 1); i >= 0; i-- {
		// fmt.Println(batteryNodes)
		for j := range batteryNodes {
			if bank[i] >= batteryNodes[j].Value {
				temp := batteryNodes[j].Value  
				tempPos := batteryNodes[j].Position
				if j-1 >= 0 {
					if batteryNodes[j-1].Position >= i {
						break
					} 
				}
				batteryNodes[j].Value = bank[i]
				batteryNodes[j].Position = i
				shiftBattery(batteryNodes[:], temp, tempPos, j+1)
				break
			}
		}
	}
	val := getCurrentJoltageArr(batteryNodes)
	// fmt.Println(val)
	return val
}

func shiftBattery(batteryNodes []BatteryNode, next int, nextPos int, j int) {
	// fmt.Println(batteryNodes[j])
	temp := batteryNodes[j].Value  
	tempPos := batteryNodes[j].Position
	if batteryNodes[j].Value <= next {
		batteryNodes[j].Value = next
		batteryNodes[j].Position = nextPos
		if j+1 < len(batteryNodes) {
			shiftBattery(batteryNodes, temp, tempPos, j+1)
		}
	}
}

func solve1(input [][]int) int {
	total := 0
	for _, bank := range input {
		total += getBatteryValue(bank)
	}
	return total
}

func solve2(input [][]int) int {
	total := 0
	for _, bank := range input {
		total += getBatteryValuePt2(bank)
	}
	return total
}

func parseFile() [][]int {
	file, err := os.Open("./part1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	arrays := [][]int{}
	for _, line := range lines {
		arrays = append(arrays, parseLargeStringIntoArray(line))
	}
	return arrays
}

func parseLargeStringIntoArray(strLineInput string) []int {
	array := make([]int, 0, len(strLineInput))

	for _, char := range strLineInput {
		digit, _ := strconv.Atoi(string(char))
		array = append(array, digit)
	}
	return array
}

func main() {
	lines := parseFile()
	result := solve2(lines)
	fmt.Println(result)
}
