// You can edit this code!
// Click here and start typing.
package main

import "fmt"

/*
Time Complexity: O(N log(N)),  Merge Sort is a recursive algorithm and time complexity can be expressed as following recurrence relation. 


*/
func merge(a []int, b []int) []int {
	var final []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	return merge(mergeSort(items[:len(items)/2]), mergeSort(items[len(items)/2:]))
}

func main() {
	arr := []int{10, 80, 30, 90, 40}
	arr = mergeSort(arr)
	fmt.Println(arr)
}

