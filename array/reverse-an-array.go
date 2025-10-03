package main

import "fmt"

func reverseArray(arr []int) {
	if len(arr) == 0 {
		return
	}

	// ****** reverse whole array ******	
	s := 0
	e := len(arr) - 1

	for s <= e {
		temp := arr[s]
		arr[s] = arr[e]
		arr[e] = temp
		s++
		e--
	}

	// reverse after a specific index m := 1,2,5,7.....
	// m := 2
	// s := m + 1
	// e := len(arr) - 1

	// for s <= e {
	// 	temp := arr[s]
	// 	arr[s] = arr[e]
	// 	arr[e] = temp
	// 	s++
	// 	e--
	// }
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	reverseArray(arr1)
	fmt.Println("reverse array: ", arr1)
}
