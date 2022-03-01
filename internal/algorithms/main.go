package algorithms

func InsertionSort(array []int) []int {

	for i := 1; i < len(array); i++ {
		x := array[i]
		j := i
		for ; j >= 1 && array[j-1] > x; j-- {
			array[j] = array[j-1]
		}
		array[j] = x
	}

	return array
}
