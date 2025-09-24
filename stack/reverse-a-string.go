package main

import "fmt"

func main() {
	s := "hello"
	stack := []rune{}
	for _, v := range s {
		stack = append(stack, v)
	}

	reversed := ""
	for len(stack) != 0 {
		reversed += string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	fmt.Println(reversed)
}