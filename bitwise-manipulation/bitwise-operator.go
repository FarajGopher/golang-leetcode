package main

import (
	"fmt"
)

func main() {

	a := 6 // 0110 in binary
	b := 5	 // 0101 in binary

	fmt.Println("a AND b: ", a & b) // 0100 in binary
	fmt.Println("a OR b: ", a|b)
	fmt.Println("a NOT b: ", ^b) 	// 1010 in binary
	fmt.Println("a XOR b: ", a^b)
	fmt.Println("a Left Shift by 1: ", a<<1) // 1100 in binary
	fmt.Println("a Right Shift by 1: ", a>>1) // 0011 in binary
}