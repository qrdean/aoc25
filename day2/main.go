package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getInvalidIds(begin int, end int) []int {
	invalidIds := []int{}
	for i := begin; i <= end; i++ {
		iString := strconv.Itoa(i)
		length := len(iString)
		if length%2 == 0 {
			half := length / 2
			first := iString[0:half]
			second := iString[half:]
			if first == second {
				invalidIds = append(invalidIds, i)
			}
		}
	}
	return invalidIds
}

func getInvalidIdsPt2(begin int, end int) []int {
	invalidIds := []int{}
	for i := begin; i <= end; i++ {
		iString := strconv.Itoa(i)
		iStringLen := len(iString)
		for length := 1; length <= (iStringLen / 2); length++ {
			// Want to get min length of pattern
			if iStringLen%length != 0 {
				continue
			}

			pattern := iString[:length]
			ok := true
			// If pattern matches then we append
			for i := length; i <= iStringLen; i += length {
				if i+length <= iStringLen {
					if iString[i:i+length] != pattern {
						ok = false
						break
					}
				}
			}

			if ok && !slices.Contains(invalidIds, i) {
				invalidIds = append(invalidIds, i)
			}
		}
	}
	return invalidIds
}

type Range struct {
	Begin int
	End   int
}

func splitRanges(input string) []Range {
	allRanges := strings.Split(input, ",")
	ranges := []Range{}
	for _, x := range allRanges {
		r := strings.Split(x, "-")
		begin, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		ranges = append(ranges, Range{Begin: begin, End: end})
	}
	return ranges
}

func transformation(input string) int {
	allRanges := strings.Split(input, ",")
	allInvalidIds := []int{}
	for _, x := range allRanges {
		r := strings.Split(x, "-")
		begin, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		invalidIds := getInvalidIds(begin, end)
		if len(invalidIds) > 0 {
			allInvalidIds = append(allInvalidIds, invalidIds...)
		}
	}
	sum := 0
	for _, num := range allInvalidIds {
		sum += num
	}
	return sum
}

func transformation2(input string) int {
	allRanges := strings.Split(input, ",")
	allInvalidIds := []int{}
	for _, x := range allRanges {
		r := strings.Split(x, "-")
		begin, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		invalidIds := getInvalidIdsPt2(begin, end)
		if len(invalidIds) > 0 {
			allInvalidIds = append(allInvalidIds, invalidIds...)
		}
	}
	sum := 0
	for _, num := range allInvalidIds {
		sum += num
	}
	return sum
}

func main() {
	input := "749639-858415,65630137-65704528,10662-29791,1-17,9897536-10087630,1239-2285,1380136-1595466,8238934-8372812,211440-256482,623-1205,102561-122442,91871983-91968838,62364163-62554867,3737324037-3737408513,9494926669-9494965937,9939271919-9939349036,83764103-83929201,24784655-24849904,166-605,991665-1015125,262373-399735,557161-618450,937905586-937994967,71647091-71771804,8882706-9059390,2546-10476,4955694516-4955781763,47437-99032,645402-707561,27-86,97-157,894084-989884,421072-462151"
	sum := transformation2(input)
	fmt.Println(sum)
}
