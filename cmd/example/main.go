package main

import (
	"log"

	pqueue "github.com/swkwon/go-priority-queue"
)

func main() {
	log.Println("first example")
	q1 := pqueue.MakePriorityQueue[int](pqueue.Low)
	q1.Push(1, 1)
	q1.Push(10, 10)
	q1.Push(2, 2)
	q1.Push(4, 4)
	q1.Push(3, 3)
	q1.Push(7, 7)
	q1.Push(8, 8)
	q1.Push(6, 6)
	for {
		v, err := q1.Pop()
		if err != nil {
			break
		}
		log.Println(v)
	}

	log.Println("second example")
	q2 := pqueue.MakePriorityQueue[int](pqueue.High)
	q2.Push(1, 1)
	q2.Push(10, 10)
	q2.Push(2, 2)
	q2.Push(4, 4)
	q2.Push(3, 3)
	q2.Push(7, 7)
	q2.Push(8, 8)
	q2.Push(6, 6)
	for {
		v, err := q2.Pop()
		if err != nil {
			break
		}
		log.Println(v)
	}
}
