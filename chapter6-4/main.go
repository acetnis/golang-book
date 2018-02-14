package main

import "fmt"

func main() {
	fmt.Println(add(1,2,3))

	xs := []int{1,2,3}
	fmt.Println(add(xs...))

	for k,v := range xs {
		fmt.Printf("k: %v, v: %v\n", k, v)
		fmt.Println("v: %v\n", v)
	}
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}