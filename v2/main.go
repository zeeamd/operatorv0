package main

// import "fmt"

import (
	"fmt"
	"sync"
)

func main() {

	// CSP - communicating sequential processes
	// channel
	// goroutine - functions
	// series of goroutine communicate with each other via channels
	// all this managed by scheduler (part of go runtime)
	// waitgroup - counter special behaviour when value = 0 (can increment/decrement)
	// Add() - increment, Done() - decrement, Wait() - wait till counter is 0

	// without this program will not wait and exit before goroutine even gets executed
	var wg sync.WaitGroup

	// channels are only type in go which must be created using built in make func
	// pass type of msg
	ch := make(chan int)

	// registering 1 go routine so add 1 - level of concurrency
	wg.Add(1)

	go func() {
		fmt.Println("async")

		// send msg to channel
		ch <- 42
	}()

	fmt.Println("sync")

	go func() {
		// recieve msg from channel
		fmt.Println(<-ch)

		// decrement counter when task complete
		wg.Done()
	}()

	// wait till counter gets to 0 - used Done()
	wg.Wait()

	// like switch works with vars select works with channel operations
	// in switch first valid case is executed top - down
	// in select it is random if more than 1 case can be executed

	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "ch1"
	}()

	go func() {
		ch2 <- "ch2"
	}()

	select {
	case msg := <- ch1:
		fmt.Println(msg)
	case msg := <- ch2:
		fmt.Println(msg)
	}

	// loop channels
	ch3 := make(chan int)
	go func() {
		for i:=0; i<5; i++ {
			ch3 <- i
		} 
		// let goroutine know no more msgs will come in
		close(ch3)
	}()

	for msg := range ch3 {
		fmt.Println(msg)
	}
}