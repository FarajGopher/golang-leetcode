package main

import "fmt"

func sortedSquares(nums []int) []int {
	i := 0
	j := len(nums) - 1
	pos := len(nums) - 1
	m := make([]int, len(nums))
	for i <= j {
		if AbsInt(nums[i]) > AbsInt(nums[j]) {
			m[pos] = nums[i] * nums[i]
			i++
		} else {
			m[pos] = nums[j] * nums[j]
			j--
		}
		pos--
	}

	return m
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	arr := []int{-4,-1,0,3,10}
	fmt.Println("sorted array: ",sortedSquares(arr))
}