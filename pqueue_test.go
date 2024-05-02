package pqueue

import (
	"log"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := MakePriorityQueue[int](Low)
	q.Push(10, 10)
	q.Push(4, 4)
	q.Push(11, 11)
	q.Push(1, 1)
	q.Push(3, 3)
	q.Push(2, 2)
	for {
		v, err := q.Pop()
		if err != nil {
			break
		} else {
			log.Println(v)
		}
	}

	type Person struct {
		Name string
		Age  int
	}

	q2 := MakePriorityQueue[*Person](High)
	q2.Push(&Person{Name: "Alice", Age: 20}, 20)
	q2.Push(&Person{Name: "Edward", Age: 10}, 10)
	q2.Push(&Person{Name: "Jessica", Age: 14}, 14)
	q2.Push(&Person{Name: "Julia", Age: 30}, 30)
	q2.Push(&Person{Name: "James", Age: 50}, 50)
	q2.Push(&Person{Name: "John", Age: 33}, 33)
	q2.Push(&Person{Name: "Argatha", Age: 13}, 13)
	q2.Push(&Person{Name: "Silva", Age: 23}, 23)
	for {
		v, err := q2.Pop()
		if err != nil {
			break
		} else {
			log.Println(v)
		}
	}
}
