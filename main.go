package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
		case t := <-time.After(time.Millisecond * 100):
			fmt.Println("timed out at", t)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		select {
		case data := <-c:
			fmt.Printf("worker %d got %d\n", w.id, data)
		case <-time.After(time.Millisecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Second)
		}
	}
}
