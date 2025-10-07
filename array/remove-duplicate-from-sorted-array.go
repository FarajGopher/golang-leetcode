package main

import (
	"fmt"
	"sort"
)

// ------------------------------------------------------------
// Approach 1: Brute Force using Map
// Time Complexity: O(n)
// Space Complexity: O(n)
// Works for: Unsorted arrays
// ------------------------------------------------------------
func removeDuplicatesBrute(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// ------------------------------------------------------------
// Approach 2: Optimal (Two Pointer Method)
// Time Complexity: O(n)
// Space Complexity: O(1)
// Works for: Sorted arrays
// ------------------------------------------------------------
func removeDuplicatesTwoPointer(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	j := 0 // index of the last unique element

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		}
	}
	return arr[:j+1]
}

func main() {
	// Example array (unsorted)
	arr := []int{3, 1, 2, 3, 4, 1, 5, 2}

	// Call Brute Force method
	fmt.Println("Original Array:", arr)
	fmt.Println("After removing duplicates (Brute Force):", removeDuplicatesBrute(arr))

	// Sort the array for two-pointer method (since it works only on sorted arrays)
	sort.Ints(arr)
	fmt.Println("\nSorted Array for Two-Pointer Method:", arr)
	fmt.Println("After removing duplicates (Two Pointer):", removeDuplicatesTwoPointer(arr))
}
