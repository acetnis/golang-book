package main

import "fmt"

func main() {
	array := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int) // channel created

	go sum(array[:len(array)/2], ch) // sum the first 3 value
	go sum(array[len(array)/2:], ch) // sum the last 3 value
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)
}

func sum(array []int, ch chan int) {
	sum := 0
	for _, value := range array {
		sum += value
	}
	ch <- sum
}
