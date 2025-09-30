package main

import "fmt"

// ***** below code can be used to find total number of item occurence 
// how ( last index - first index ) + 1 


//when mid found
//max = mid - 1

func firstOccurence(arr []int,find int) int{
	min := 0
	max := len(arr) - 1
	first := -1
	for min <= max {
		mid := min + ((max - min) / 2)
		if arr[mid] == find {
			first = mid
			max = mid - 1
		} else if arr[mid] < find {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return first
}

//when mid found
//min = mid + 1

func lastOccurence(arr []int,find int) int{
	min := 0
	max := len(arr) - 1
	last := -1
	for min <= max {
		// mid := (min + max) / 2
		mid := min + ((max - min) / 2)
		if arr[mid] == find {
			last = mid
			min = mid + 1
		} else if arr[mid] < find {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return last
}

func main() {
	arr := []int{0, 0, 1, 1, 2, 2, 2, 2}
	find := 2
	first := firstOccurence(arr,find)
	last := lastOccurence(arr,find)

	fmt.Println("first and last occurence:",first,last)
}
