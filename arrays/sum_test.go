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
	t.Run("calculates the sum of 5 numbers", func(t *testing.T) {
		sliceOne := []int{1, 2}
		sliceTwo := []int{0, 9}

		got := SumAll(sliceOne, sliceTwo)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v, given %v and %v", got, want, sliceOne, sliceTwo)
		}

	})
}
