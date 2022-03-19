package algorithms

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {

	SORTED := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	UNSORTED := []int{2, 3, 5, 1, 4, 10, 8, 7, 6, 9, 0}

	InsertionSort(UNSORTED)

	for i, a := range UNSORTED {
		if a != SORTED[i] {
			t.Error("Insertion sort: expected ", SORTED[i], " got ", a)
		}
	}
}
