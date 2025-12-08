package main

import "testing"

func TestMain(t *testing.T) {
	// testInput := "987654321111111\n811111111111119\n234234234234278\n818181911112111\n"
	// t.Run("test", func(t *testing.T) {
	// 	got := solve1(testInput)
	// 	want := 357
	// 	if got != want {
	// 		t.Errorf("got: %d, want:%d", got, want)
	// 	}
	// })

	t.Run("get highest two nodes", func(t *testing.T) {
		highest, secondHighest := getHighestTwo([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1})
		if highest.Value != 9 {
			t.Errorf("highest value is not 9 but %d", highest.Position)
		}
		if highest.Position != 0 {
			t.Errorf("highest pos is not 0 but %d", highest.Position)
		}

		if secondHighest.Value != 8 {
			t.Errorf("secondHighest value is not 9 but %d", secondHighest.Position)
		}
		if secondHighest.Position != 1 {
			t.Errorf("secondHighest pos is not 0 but %d", secondHighest.Position)
		}

		highest, secondHighest = getHighestTwo([]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9})
		if highest.Value != 9 {
			t.Errorf("highest value is not 9 but %d", highest.Position)
		}
		if highest.Position != 14 {
			t.Errorf("highest pos is not 15 but %d", highest.Position)
		}

		if secondHighest.Value != 8 {
			t.Errorf("secondHighest value is not 8 but %d", secondHighest.Position)
		}
		if secondHighest.Position != 0 {
			t.Errorf("secondHighest pos is not 0 but %d", secondHighest.Position)
		}

		highest, secondHighest = getHighestTwo([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8})
		if highest.Value != 8 {
			t.Errorf("highest value is not 9 but %d", highest.Position)
		}
		if highest.Position != 14 {
			t.Errorf("highest pos is not 15 but %d", highest.Position)
		}

		if secondHighest.Value != 7 {
			t.Errorf("secondHighest value is not 8 but %d", secondHighest.Position)
		}
		if secondHighest.Position != 13 {
			t.Errorf("secondHighest pos is not 0 but %d", secondHighest.Position)
		}

		highest, secondHighest = getHighestTwo([]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1})
		if highest.Value != 9 {
			t.Errorf("highest value is not 9 but %d", highest.Position)
		}
		if highest.Position != 6 {
			t.Errorf("highest pos is not 6 but %d", highest.Position)
		}

		if secondHighest.Value != 1 {
			t.Errorf("secondHighest value is not 8 but %d", secondHighest.Position)
		}
		if secondHighest.Position != 0 {
			t.Errorf("secondHighest pos is not 0 but %d", secondHighest.Position)
		}
	})
}
