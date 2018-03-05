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
	fbTemplate := func(fbnumber int, str string) func(n int) (string, bool) {
		return func(n int) (string, bool) {
			if n%fbnumber == 0 {
				return str, true
			}
			return "", false
		}
	}
	fbArray := [...]func(n int) (string, bool){
		fbTemplate(15, "FizzBuzz"),
		fbTemplate(3, "Fizz"),
		fbTemplate(5, "Buzz"),
	}

	for i := 0; i < len(fbArray); i++ {
		if str, ok := fbArray[i](number); ok {
			return str
		}
	}

	return strconv.Itoa(number)
}
