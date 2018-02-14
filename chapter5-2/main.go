package main

import "fmt"

/*
func main() {
	for number := 1; number <= 100; number++ {
		if number%15 == 0 {
			fmt.Println(number, "FizzBuzz")
		} else if number%3 == 0 {
			fmt.Println(number, "Fizz")
		} else if number%5 == 0 {
			fmt.Println(number, "Buzz")
		} else {
			fmt.Println(number)
		}
	}
}
*/

func main() {
	number := []int{100,15,3,5}
	fizzbuzz(number[0],number[1],number[2],number[3])
}

func fizzbuzz(round int,x1 int,x2 int,x3 int) {
	for number := 1; number <= round; number++ {
		if number%x1 == 0 {
			print(number, "FizzBuzz")
		} else if number%x2 == 0 {
			print(number, "Fizz")
		} else if number%x3 == 0 {
			print(number, "Buzz")
		} else {
			fmt.Println(number)
		}
	}
}

func print(n int,str string) {
	fmt.Println(n, str)
}