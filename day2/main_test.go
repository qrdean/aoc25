package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestMain(t *testing.T) {
	testInput := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	t.Run("result test", func(t *testing.T) {
		got := transformation(testInput)
		want := 1227775554
		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("result test 2", func(t *testing.T) {
		got := transformation2(testInput)
		want := 4174379265
		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("test range splitting", func(t *testing.T) {
		ranges := splitRanges(testInput)
		wantRanges := []Range{
			{Begin: 11, End: 22},
			{Begin: 95, End: 115},
			{Begin: 998, End: 1012},
			{Begin: 1188511880, End: 1188511890},
			{Begin: 222220, End: 222224},
			{Begin: 1698522, End: 1698528},
			{Begin: 446443, End: 446449},
			{Begin: 38593856, End: 38593862},
			{Begin: 565653, End: 565659},
			{Begin: 824824821, End: 824824827},
			{Begin: 2121212118, End: 2121212124},
		}
		for _, want := range wantRanges {
			if !slices.Contains(ranges, want) {
				t.Errorf("ranges: %v does not contain: %v", ranges, want)
			}
		}
	})

	t.Run("test getting numbers", func(t *testing.T) {
		got := getInvalidIds(11, 22)
		want := []int{11, 22}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIds(1188511880, 1188511890)
		want = []int{1188511885}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIds(38593856, 38593862)
		want = []int{38593859}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIds(998, 1012)
		want = []int{1010}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIds(222220, 222224)
		want = []int{222222}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIds(1698522, 1698528)
		want = []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIds(446443, 446449)
		want = []int{446446}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIds(446443, 446449)
		want = []int{446446}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test invalid numbers pt2", func(t *testing.T) {
		got := getInvalidIdsPt2(11, 22)
		want := []int{11, 22}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(95, 115)
		want = []int{99, 111}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(998, 1012)
		want = []int{999, 1010}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(1188511880, 1188511890)
		want = []int{1188511885}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(222220, 222224)
		want = []int{222222}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(38593856, 38593862)
		want = []int{38593859}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(1698522, 1698528)
		want = []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(446443, 446449)
		want = []int{446446}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(824824821, 824824827)
		want = []int{824824824}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = getInvalidIdsPt2(2121212118, 2121212124)
		want = []int{2121212121}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func BenchmarkPt1(b *testing.B) {
	input := "749639-858415,65630137-65704528,10662-29791,1-17,9897536-10087630,1239-2285,1380136-1595466,8238934-8372812,211440-256482,623-1205,102561-122442,91871983-91968838,62364163-62554867,3737324037-3737408513,9494926669-9494965937,9939271919-9939349036,83764103-83929201,24784655-24849904,166-605,991665-1015125,262373-399735,557161-618450,937905586-937994967,71647091-71771804,8882706-9059390,2546-10476,4955694516-4955781763,47437-99032,645402-707561,27-86,97-157,894084-989884,421072-462151"
	transformation(input)
}

func BenchmarkPt2(b *testing.B) {
	input := "749639-858415,65630137-65704528,10662-29791,1-17,9897536-10087630,1239-2285,1380136-1595466,8238934-8372812,211440-256482,623-1205,102561-122442,91871983-91968838,62364163-62554867,3737324037-3737408513,9494926669-9494965937,9939271919-9939349036,83764103-83929201,24784655-24849904,166-605,991665-1015125,262373-399735,557161-618450,937905586-937994967,71647091-71771804,8882706-9059390,2546-10476,4955694516-4955781763,47437-99032,645402-707561,27-86,97-157,894084-989884,421072-462151"
	transformation2(input)
}
