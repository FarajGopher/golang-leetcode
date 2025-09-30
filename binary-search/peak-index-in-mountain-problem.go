package main

import "fmt"

func findPivot(arr []int) int {
	start, end := 0, len(arr)-1

	for start < end {
		mid := (start + end) / 2
		if arr[mid] < arr[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func main() {
	arr := []int{0, 2, 4, 6, 5, 3, 1}
	fmt.Println("Peak index:", findPivot(arr))  
}
