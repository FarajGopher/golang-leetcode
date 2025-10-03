package main

import "fmt"

func sortTwoArray(arr1 []int, len1 int, arr2 []int, len2 int) []int {
	i := 0
	j := 0
	newarr := []int{}
	for i < len1 && j < len2 {
		if arr1[i] < arr2[j] {
			newarr = append(newarr, arr1[i])
			i++
		} else {
			newarr = append(newarr, arr2[j])
			j++
		}
	}

	// Append remaining elements of arr1 (if any)
	for i < len1 {
		newarr = append(newarr, arr1[i])
		i++
	}

	// Append remaining elements of arr2 (if any)
	for j < len2 {
		newarr = append(newarr, arr2[j])
		j++
	}

	return newarr
}

func main() {
	arr1 := []int{1, 2, 5, 7, 8}
	arr2 := []int{3, 4, 6, 9}

	newSortedArray := sortTwoArray(arr1, len(arr1), arr2, len(arr2))
	fmt.Println("new sorted array after merge: ", newSortedArray)
}
