package main

import "fmt"

func checkArraySorted(arr []int) bool{
	for i :=0 ;i<len(arr)-1;i++ {
		if arr[i]>arr[i+1]{
			return false
		}
	}
	
	return true
}

func main() {
	arr := []int{1, 2, 12, 4, 5}

	check := checkArraySorted(arr)
	fmt.Println("check status:",check)
}
