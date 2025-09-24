package stack

func MinCostBracketValid(s string) int{
	//if cost is odd then we dnn't need to transform as we can never make all valid pairs
	if len(s)%2 == 1{
		return -1
	}

	stack := []rune{}

	for _,v := range s {
		if v == '{'{
			stack = append(stack, v)
		}else { //means we encounter closing bracket so we need to look for top of the stack if its ( open bracket then pop it casue we find the pair
			if len(stack) != 0 && stack[len(stack)-1] == '{'{
				stack = stack[:len(stack)-1]
			}else{ // else we need to put the element in stack
				stack = append(stack, v)
			}
		}
	}

	a ,b := 0,0

	for len(stack) == 0 {
		if stack[len(stack)-1] == '{'{
			b++
		}else{
			a++
		}
		stack = stack[:len(stack)-1]
	}

	ans := (a+1)/2 + (b+1)/2


	return ans
}