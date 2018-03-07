package main

import "fmt"

func main() {
	c := make(chan int)
	//c <- 1
	//c <- 2
	go func() { c <- 1 }()
	go func() { c <- 2 }()
	fmt.Println(<-c)
	fmt.Println(<-c)
}

/*
A send operation on an unbuffered channel blocks the sending
goroutine until another goroutine executes a corresponding
reveive on the same channel, at which point the value is
transmitted and both goroutine may continue
*/
