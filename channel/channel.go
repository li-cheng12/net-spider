package main

import (
	"fmt"
	"sync"
)

type worker2 struct {
	ch chan int
	done func()
}

type worker struct {
	ch chan int
	done chan bool
}
func doWork(id int) worker {
	w := worker{}
	w.ch = make(chan int)
	w.done = make(chan bool)
	go func() {
		for {
			c := <-w.ch
			fmt.Printf("%c, %d\n", c, id)
			w.done <-true
		}
	}()
	return w
}

func doWork2(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		ch: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go func() {
		for {
			c := <-w.ch
			fmt.Printf("%c, %d\n", c, id)
			w.done()
		}
	}()
	return w
}

func createWork() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = doWork(i)
	}
	for i:=0; i < 10; i++ {
		workers[i].ch <- 'a' + i
	}
	for i := 0;i < 10; i++ {
		<- workers[i].done
	}
	for i:=0; i < 10; i++ {
		workers[i].ch <- 'A' + i
	}
	for i := 0;i < 10; i++ {
		<- workers[i].done
	}

}

func createWorker2 () {
	var wg sync.WaitGroup
	var workers [10]worker2
	for i := 0; i < 10; i++ {
		workers[i] = doWork2(i, &wg)
	}
	wg.Add(10)
	for i, w := range workers {
		w.ch <- 'a' + i
	}
	wg.Wait()

}
func main() {
	createWorker2()
}
