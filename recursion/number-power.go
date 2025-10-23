package main

import "fmt"

func multiplier(v, n int) int {
	if n == 0 {
		return 1
	}
	return v * multiplier(v, n-1)
}

func main() {
	v := multiplier(2, 5)
	fmt.Println(v)
}
