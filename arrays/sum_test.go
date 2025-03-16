package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("calculates the sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}

	})

	t.Run("calculates the sum of slice of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}

	})

}

func TestSumAll(t *testing.T) {
	t.Run("calculates the sum all slices", func(t *testing.T) {
		sliceOne := []int{1, 2}
		sliceTwo := []int{0, 9}

		got := SumAll(sliceOne, sliceTwo)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v, given %v and %v", got, want, sliceOne, sliceTwo)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("calculates the sum of all but the first number in the slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9, 1})
		want := []int{2, 10}

		checkSums(t, got, want)
	})

	t.Run("Sum of empty slice should return 0", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}

		checkSums(t, got, want)
	})

}
