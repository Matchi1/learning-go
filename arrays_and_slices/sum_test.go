package main

import (
    "testing"
    "reflect"
)

func TestSum(t *testing.T) {
    t.Run("Collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}

        got := Sum(numbers)
        want := 15

        if got != want {
            t.Errorf("Expected %d but got %d given, %v", want, got, numbers)
        }
    })
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
