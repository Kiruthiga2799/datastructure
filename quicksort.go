package main

import (
	"fmt"
)

func main() {
	data := []int{90, 19, 87, 27, 71, 34, 88, 78, 56}
	QuickSort(data)
	fmt.Println(data)
}

func less(value1 int, value2 int) bool {
	return value1 < value2
}

func more(value1 int, value2 int) bool {
	return value1 > value2
}

func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func QuickSort(arr []int) {
	size := len(arr)
	quickSortUtil(arr, 0, size-1)
}

func quickSortUtil(arr []int, lower int, upper int) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for arr[lower] <= pivot && lower < upper {
			lower++
		}
		for arr[upper] > pivot && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}
	swap(arr, upper, start)            // upper is the pivot position
	quickSortUtil(arr, start, upper-1) // pivot -1 is the upper for left sub array.
	quickSortUtil(arr, upper+1, stop)  // pivot + 1 is the lower for right sub array.
}
