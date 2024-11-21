package main

import "fmt"

// Queue
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Deqeue

func (q *Queue) Deqeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

var p = fmt.Println

type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	return toRemove
}

func main() {
	myStack := Stack{}
	p(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.Pop()
	p(myStack)
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Deqeue()

	p("main execution started")
	p("No. of CPUs:", runtime.NumCPU())
	p("No. of Goroutines:", runtime.NumGoroutine())

	p("OS:", runtime.GOOS)
	p("Arch:", runtime.GOARCH)

	p("GOMAXPROCS", runtime.GOMAXPROCS(0))

}
