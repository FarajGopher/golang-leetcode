package main

import "fmt"

// func main() {
// 	//queue
// 	a := []int{-8, 2, 3, -6, 10}

// 	//window size
// 	k := 2
// 	result := []int{}
	


// 	for i := 0; i <= len(a)-k; i++ {
// 		firstNeg := 0
// 		for j := i; j < i+k ; j++ {
// 			if a[j] < 0 {
// 				firstNeg = a[j]
// 				break
// 			}
// 		}
// 		result = append(result, firstNeg)
// 	}

// 	fmt.Println("new array : ",result)
// }


// ******** Using Doubly ended queue *********

type Deque struct {
	arr        []int
	front, rear int
	size        int
}

// Pushes 'x' at the rear
func (d *Deque) PushRear(x int) bool {
	if d.IsFull() {
		return false
	} else if d.IsEmpty() {
		d.front, d.rear = 0, 0
	} else if d.rear == d.size-1 && d.front != 0 {
		d.rear = 0
	} else {
		d.rear++
	}
	d.arr[d.rear] = x
	return true
}


// Check if empty
func (d *Deque) IsEmpty() bool {
	return d.front == -1
}

// Check if full
func (d *Deque) IsFull() bool {
	return (d.front == 0 && d.rear == d.size-1) ||
		(d.rear == (d.front-1+d.size)%d.size)
}

func main(){
	size := 10
	k:= 2
	d := Deque{
		size:  10,
		front: -1,
		rear:  -1,
		arr:   make([]int, size),
	}

	for i := 0;i<k;i++{
		if d.arr[i]<0 {
			d.PushRear(i)
		}
	}

	if d.size > 0{

	}

}