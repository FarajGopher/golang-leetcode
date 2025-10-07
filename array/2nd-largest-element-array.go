package main

import (
	"fmt"
	"math"
)

//**** better approach time complexity O(n+n) = O(2n)
func secondlargestElementBetter(arr []int) int {
	largest := arr[0]
	secondLar := -1
	for i := 1; i < len(arr); i++ {
		if largest < arr[i] {
			largest = arr[i]
		}
	}
	for i := 0; i < len(arr); i++ {
		if secondLar < arr[i] && arr[i] != largest {
			secondLar = arr[i]
		}
	}
	return secondLar
}

// ******* optimal approach O(N)
func secondlargestElementOptimal(arr []int) int {
	largest := arr[0]
	secondLar := -1
	for i := 1; i < len(arr); i++ {
		if largest < arr[i]{
			secondLar = largest
			largest = arr[i]
		}else if secondLar < arr[i] && arr[i] != largest {
			secondLar = arr[i]
		} 
	}
	return secondLar
}

// ******* optimal approach
func secondSmallestElementOptimal(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	smallest := math.Inf(1)
	secondSmallest := math.Inf(1)

	for _, v := range arr {
		if float64(v) < smallest {
			secondSmallest = smallest
			smallest = float64(v)
		} else if float64(v) < secondSmallest && float64(v) != smallest {
			secondSmallest = float64(v)
		}
	}

	// If secondSmallest is still infinity, no valid second smallest
	if secondSmallest == math.Inf(1) {
		return int(secondSmallest)
	}

	return int(secondSmallest)
}

func main() {
	arr := []int{10, 5, 20, 8}

	fmt.Println("secondlargest better approach: ", secondlargestElementBetter(arr))
	fmt.Println("optimal approach: ",secondlargestElementBetter(arr))
	fmt.Println("second smallest element : ",secondSmallestElementOptimal(arr))
}
