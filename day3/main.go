package main

import (
	"fmt"
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

type BatteryNode struct {
	Value    int
	Position int
}

func getBatteryValue(bank []int) int {
	first := BatteryNode{Value:0, Position: 0}
	second := BatteryNode{Value:0, Position: 0}
	inCase := 0
	for i := len(bank) - 1; i >= 0; i-- {
		if second.Value < bank[i] {
			second.Value = bank[i]
			second.Position = i
		}
		if first.Value < second.Value {
			temp := first
			first = second
			second = temp		
		}
	}
}

func getHighestTwo(bank []int) (BatteryNode, BatteryNode) {
	currentJoltage := 0
	highOne := BatteryNode{Value: 0, Position: 0}
	highTwo := BatteryNode{Value: 0, Position: 0}
	// bankNode := BankNode{First: &highOne, Second: &highTwo}
	for i, x := range bank {
		if highTwo.Value < x {
			highTwo.Value = x
			highTwo.Position = i
		}
		if (getCurrentJoltage(highOne.Value, highTwo.Value) > currentJoltage) && highOne.Value < highTwo.Value {
			temp := highOne.Value
			tempPos := highOne.Position
			highOne.Value = highTwo.Value
			highOne.Position = highTwo.Position
			highTwo.Value = temp
			highTwo.Position = tempPos
		}

		if getCurrentJoltage(highOne.Value, highTwo.Value) > currentJoltage {
			currentJoltage = getCurrentJoltage(highOne.Value, highTwo.Value)
		}
	}
	fmt.Println(currentJoltage)
	return highOne, highTwo
}

func solve1(input string) int {
	return 0
}

func main() {
}
