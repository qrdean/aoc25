package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("check battery values", func(t *testing.T) {

		batteryValue := getBatteryValue([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1})
		want := 98
		if batteryValue != want {
			t.Errorf("batteryValue: %d is not %d", batteryValue, want)
		}
		batteryValue= getBatteryValue([]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9})
		want = 89
		if batteryValue != want {
			t.Errorf("batteryValue: %d is not %d", batteryValue, want)
		}
		batteryValue= getBatteryValue([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8})
		want = 78
		if batteryValue != want {
			t.Errorf("batteryValue: %d is not %d", batteryValue, want)
		}
		batteryValue= getBatteryValue([]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1})
		want = 92
		if batteryValue != want {
			t.Errorf("batteryValue: %d is not %d", batteryValue, want)
		}
	})
	
	t.Run("parse large number into array", func(t *testing.T) {
		testInput := "123456789"
		got := parseLargeStringIntoArray(testInput)
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(got, want) {
    	t.Errorf("got: %v want :%v", got, want)
		}
		
		testInput = "2555245573282137352766682525526364435746545343523394355638332326665366122245646523573255525564158774"
		got = parseLargeStringIntoArray(testInput)
		want = []int{2, 5, 5, 5, 2, 4, 5, 5, 7, 3, 2, 8, 2, 1, 3, 7, 3, 5, 2, 7, 6, 6, 6, 8, 2, 5, 2, 5, 5, 2, 6, 3, 6, 4, 4, 3, 5, 7, 4, 6, 5,4,5,3, 4,3,5,2,3,3,9,4,3,5,5,6,3,8,3,3,2,3,2,6,6,6,5,3,6,6,1,2,2,2,4,5,6,4,6,5,2,3,5,7,3,2,5,5,5,2,5,5,6,4,1,5,8,7,7,4}
		if !reflect.DeepEqual(got, want) {
    	t.Errorf("got: %v want :%v", got, want)
		}
		
	})

	t.Run("some test", func(t *testing.T) {
		testRealInput := "2555245573282137352766682525526364435746545343523394355638332326665366122245646523573255525564158774\n5413546422442523229295232262733493414148154532639862333271175557334235334296623642342223646213476455\n5433332355353372756453622653442176834335623343626343236463374522272945534432336513564562366234336223"
		testThing := strings.Split(testRealInput, "\n")
		results := []int{98, 99, 96}
		for i, line := range testThing {
			bank := parseLargeStringIntoArray(line)
			val := getBatteryValue(bank)
			if val != results[i] {
				t.Errorf("got: %d, want :%d", val, results[i])
			}
		}
	})
	
	t.Run("check battery values pt2", func(t *testing.T) {

		batteryValue := getBatteryValuePt2([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1})
		want := 987654321111
		if batteryValue != want {
			t.Errorf("batteryValue: %d is not %d", batteryValue, want)
		}
		batteryValue= getBatteryValuePt2([]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9})
		want = 811111111119
		if batteryValue != want {
			t.Errorf("batteryValue: %d is not %d", batteryValue, want)
		}
		batteryValue= getBatteryValuePt2([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8})
		want = 434234234278
		if batteryValue != want {
			t.Errorf("batteryValue: %d is not %d", batteryValue, want)
		}
		batteryValue= getBatteryValuePt2([]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1})
		want = 888911112111
		if batteryValue != want {
			t.Errorf("batteryValue: %d is not %d", batteryValue, want)
		}
	})

	t.Run("pt2 test real", func(t *testing.T) {
		testRealInput := "2555245573282137352766682525526364435746545343523394355638332326665366122245646523573255525564158774\n5413546422442523229295232262733493414148154532639862333271175557334235334296623642342223646213476455\n5433332355353372756453622653442176834335623343626343236463374522272945534432336513564562366234336223"
		testThing := strings.Split(testRealInput, "\n")
		results := []int{987564158774, 999996676455, 966666436223}
		for i, line := range testThing {
			bank := parseLargeStringIntoArray(line)
			val := getBatteryValuePt2(bank)
			if val != results[i] {
				t.Errorf("got: %d, want :%d", val, results[i])
			}
		}
	})
}

func BenchmarkPt1(b *testing.B) {
	lines := parseFile()
	b.ResetTimer()
	solve1(lines)	
}

func BenchmarkPt2(b *testing.B) {
	lines := parseFile()
	b.ResetTimer()
	solve2(lines)	
}
