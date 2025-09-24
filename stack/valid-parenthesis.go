package main

import "fmt"

func CheckValidParenthesis(s string) bool {
	stack := []rune{}
	items := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) != 0 {
				top := stack[len(stack)-1]
				if top == items[v] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}

	}

	return len(stack) == 0
}

func main() {
	s := "{{()}}"
	check := CheckValidParenthesis(s)
	fmt.Println("is valid parenthesis: ", check)
}
