package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

	fmt.Print("Enter Degree Fahrenheit: ")
	var fdegree float64
	fmt.Scanf("%f", &fdegree)
	cdegree := (fdegree - 32) * 5 / 9
	fmt.Printf("Degree Celcius is %.2f", cdegree)
}