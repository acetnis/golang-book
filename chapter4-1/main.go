package main

import "fmt"

func main() {
	const x string = "Hello, World"
	v1, v2 := "first", "sec"
	v1, v2 = v2, v1

	fmt.Println(v1)
	fmt.Println(v2)
}