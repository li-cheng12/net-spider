package main

import (
	"time"
	"math/rand"
	"fmt"
)

func Work () chan int {
	ch := make(chan int)
	 go func() {
		 for {

			 i := <- ch
			 time.Sleep(time.Second)
			 fmt.Println(i)
		 }
	 }()
	 return ch
}
func generate() chan int{
	ch := make(chan int)
	go func() {
		count := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			ch <- count
			count++
		}
	}()
	return ch
}

func main() {
	c1, c2 := generate(), generate()
	arr := []int{}
		tm := time.After(time.Second *10)
	for {
		var ch chan int
		var value int
		if len(arr) > 0 {
			ch = Work()
			value = arr[0]
		}
		select {
		case n := <- c1:
			arr = append(arr,n)
		case n := <- c2:
			arr = append(arr,n)
		case ch <- value:
			arr = arr[1:]
		case <-time.After(800*time.Millisecond):
			fmt.Println("time out")

		case <-tm:
			fmt.Println("bye")
			return
		}

	}
}
