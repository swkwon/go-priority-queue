package pqueue

import (
	"errors"
	"sync"
)

var ErrQueueIsEmpty = errors.New("queue is empty")

type OrderBy int

const (
	Low  OrderBy = 0
	High OrderBy = 1
)

type Element[T any] struct {
	priority  int
	container T
}

func (e *Element[T]) GetPriority() int {
	return e.priority
}

func (e *Element[T]) GetValue() T {
	return e.container
}

type PQueue[T any] struct {
	elements []*Element[T]
	mutex    *sync.Mutex
	order    OrderBy
}

func New[T any](popOrder OrderBy) *PQueue[T] {
	return &PQueue[T]{
		mutex: &sync.Mutex{},
		order: popOrder,
	}
}

func (p *PQueue[T]) compare(i, j int) bool {
	switch p.order {
	case Low:
		return p.elements[i].GetPriority() > p.elements[j].GetPriority()
	case High:
		return p.elements[i].GetPriority() < p.elements[j].GetPriority()
	default:
		return false
	}
}

func (p *PQueue[T]) Len() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return len(p.elements)
}

func (p *PQueue[T]) Push(v T, priority int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.elements = append(p.elements, &Element[T]{
		priority:  priority,
		container: v,
	})

	ci := len(p.elements) - 1
	for ci != 0 {
		pi := (ci - 1) / 2
		if p.compare(pi, ci) {
			temp := p.elements[pi]
			p.elements[pi] = p.elements[ci]
			p.elements[ci] = temp
			ci = pi
		} else {
			break
		}
	}
}

func (p *PQueue[T]) Pop() (T, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	var ret T
	if len(p.elements) <= 0 {
		return ret, ErrQueueIsEmpty
	}

	pi := 0
	ret = p.elements[pi].GetValue()
	lastIndex := len(p.elements) - 1
	p.elements[pi] = p.elements[lastIndex]
	p.elements = p.elements[:lastIndex]
	lastIndex--

	for {
		left, right := (pi*2)+1, (pi*2)+2
		if left > lastIndex {
			break
		}
		ci := left

		if right <= lastIndex && p.compare(left, right) {
			ci = right
		}
		if !p.compare(pi, ci) {
			break
		}

		temp := p.elements[pi]
		p.elements[pi] = p.elements[ci]
		p.elements[ci] = temp
		pi = ci
	}

	return ret, nil
}
