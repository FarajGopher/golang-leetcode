package main

import "fmt"

//prefix sum
func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	val := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		val[i] = energy[i]
		if i < n-k {
			val[i] += val[i+k]
		}
	}
	max := val[0]
	for i := 1; i < n; i++ {
		if val[i] > max {
			max = val[i]
		}
	}
	return max
}



func main() {
	fmt.Println(maximumEnergy([]int{5, 2, -10, -5, 1}, 3)) // Output: 3
	fmt.Println(maximumEnergy([]int{-2, -3, -1}, 2))       // Output: -1
	fmt.Println(maximumEnergy([]int{-5, 10, -2, 4}, 2))    // Output: 14
}