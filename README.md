# go-priority-queue
 go-priority-queue는 우선순위가 빠른 요소를 먼저 pop 할 수 있도록 구현한 thread safety한 generic 알고리즘입니다.
# 패키지 설치
```
$ go get github.com/swkwon/go-priority-queue@latest
```
# 시작하기
```
package main

import (
	"log"

	pqueue "github.com/swkwon/go-priority-queue"
)

func main() {
	log.Println("first example")
	q1 := pqueue.New[int](pqueue.Low)
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
	q2 := pqueue.New[int](pqueue.High)
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
```
`New`로 큐에 들어갈 요소의 타입과 우선순위 결정방식을 지정합니다. `Low`는 우선순위의 값이 작을 수록 `High`는 우선순위의 값이 클 수록 먼저 큐로 부터 나오게 됩니다.