package main

import "fmt"

//n=5
//n-1 * n-2 * n-3 ...

func calculateFactorial(n int)int{
	// if n == 0{
	// 	return 1
	// }
	return n * calculateFactorial(n-1)
}

func main(){
	result := calculateFactorial(5)
	fmt.Println("result: ",result)
}