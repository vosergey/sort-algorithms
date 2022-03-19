package main

import (
	"fmt"
	"sort-algorithms/internal/algorithms"
)

func main() {

	unsortedArray := []int{8, 7, 6, 5, 4, 3, 2, 1}

	sortedArray := algorithms.InsertionSort(unsortedArray)

	fmt.Println(sortedArray)
}
