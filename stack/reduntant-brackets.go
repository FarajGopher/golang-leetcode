package main

import "fmt"

//sudo
/* for range till end of expression
if ch == '(' || operator
	push in stack
else if ch == ')'
	if stack include operator (not reduntant)
		continue
	else (redundant)
		return false
*/

func IsRedundant(exp string) bool {
	stack := []rune{}
	for _,ch := range exp{
		if ch == '(' || ch == '+' || ch == '-'|| ch == '*' || ch == '/' {
			stack = append(stack, ch)
		}else{
			if ch == ')' {
				isRedundant := true
				for stack[len(stack)-1] != '(' {
					top := stack[len(stack)-1]
					if top == '+' || top == '-'|| top == '*' || top == '/'{
						isRedundant = false
					}
					stack = stack[:len(stack)-1]
				}
				if isRedundant{
					return true
				}

				stack = stack[:len(stack)-1]
			}
		}
	}
	return false
}


func main() {
	// expression := "((a*b)+(c/d))"
	redundantExpression := "(/)"
	check := IsRedundant(redundantExpression)
	fmt.Println("is reduntant: ", check)
}


