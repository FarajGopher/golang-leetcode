package main

import (
	"fmt"
	"leetcode/stack"
)

func main() {
	// expression := "((a*b)+(c/d))"
	redundantExpression := "(/)"
	check := stack.IsRedundant(redundantExpression)
	fmt.Println("is reduntant: ", check)

	cost := stack.MinCostBracketValid("{{}}}}")
	fmt.Println("cost:",cost)

}