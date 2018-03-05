package main

import (
	"fmt"
	"strconv"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(number, fizzbuzz(number))
	}
}

func fizzbuzz(number int) string {
	fizzBuzzFunc := func(n int) (string, bool) {
		if n%15 == 0 {
			return "FizzBuzz", true
		}
		return "", false
	}
	fizzFunc := func(n int) (string, bool) {
		if n%3 == 0 {
			return "Fizz", true
		}
		return "", false
	}
	buzzFunc := func(n int) (string, bool) {
		if n%5 == 0 {
			return "buzz", true
		}
		return "", false
	}
	fbArray := [...]func(n int) (string, bool){
		fizzBuzzFunc,
		buzzFunc,
		fizzFunc,
	}

	for i := 0; i < len(fbArray); i++ {
		if str, ok := fbArray[i](number); ok {
			return str
		}
	}

	return strconv.Itoa(number)
}
