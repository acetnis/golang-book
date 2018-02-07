package main

import "fmt"

func main() {
	fmt.Println("=====Floating point=====")
	third := 1.0 / 3.0
	fmt.Println("third = %v\n", third)
	fmt.Printf("third + third + third = %v\n", third+third+third)
}