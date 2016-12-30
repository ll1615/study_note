package main

import "fmt"

func main() {
	var sp semaphore
	sp = make(chan Empty)
	go func(ch chan Empty) {
		for i := 0; i < 10; i++ {
			ch <- i * 10
		}
	}(sp)
	sp.V(10)
}

type Empty interface{}
type semaphore chan Empty

func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resouces
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s //fmt.Println(<-s)
	}
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
