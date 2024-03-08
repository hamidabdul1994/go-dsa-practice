package main

import "fmt"


/*
Best Case: Ω (N log (N))
The best-case scenario for quicksort occur when the pivot chosen at the each step divides the array into roughly equal halves.
In this case, the algorithm will make balanced partitions, leading to efficient Sorting.
Average Case: θ ( N log (N))
Quicksort’s average-case performance is usually very good in practice, making it one of the fastest sorting Algorithm.
Worst Case: O(N2)
*/

// Put Pivot value at correct position
func partition(arr []int, low, high int) int {
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < arr[high] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// QuickSort: Get Pivot and place in Correct possition Also divide the Worker to do sort 
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 80, 30, 90, 40}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

