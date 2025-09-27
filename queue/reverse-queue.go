package main

import "fmt"

type Queue struct {
	arr   []int
	front int
	rear  int
	size  int
}

type Stack struct {
	item []int
}

// Constructor
func NewQueue(size int) *Queue {
	return &Queue{
		size:  size,
		arr:   make([]int, size),
		front: 0,
		rear:  0,
	}
}

// Enqueue operation
func (q *Queue) Enqueue(x int) {
	if q.rear == q.size {
		fmt.Println("cann't insert queue is full")
		return
	}
	q.arr[q.rear] = x
	q.rear++
}

func (q *Queue) Dequeue() int{
	if q.front == q.rear {
		fmt.Println("Queue is empty cannot dequeue")
		return -1
	}
	val:= q.arr[q.front]
	q.front++
	if q.front == q.rear{
		q.front = 0
		q.rear = 0
	}
	return val
}

func (q *Queue) IsEmpty() bool{
	return q.front == q.rear 
}

func main() {
	stack := Stack{
		item: nil,
	}
	q := NewQueue(10)
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("queue before reverse: ",q.arr)

	for !q.IsEmpty() {
		val := q.Dequeue()
		stack.item = append(stack.item, val)
	}

	for len(stack.item) != 0 {
		top := stack.item[len(stack.item)-1]
		q.Enqueue(top)
		stack.item = stack.item[:len(stack.item)-1]
	}

	fmt.Println("queue after reverse: ",q.arr)

}



