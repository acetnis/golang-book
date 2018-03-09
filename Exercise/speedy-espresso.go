package main

import (
	"fmt"
	"time"
)

func main() {
	volumn := 200
	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func order(volumn int) (container []string) {
	for i := 1; i <= volumn; i++ {
		// cashier receive order
		time.Sleep(5 * time.Millisecond)
		coffee := fmt.Sprintf("order: %d", i)

		// barista brew coffee
		//time.Sleep(100 * time.Millisecond)
		//coffee = fmt.Sprintf("%s %s", coffee, "espresso")

		go barista(coffee)

		// waiter serve coffee
		time.Sleep(5 * time.Millisecond)
		container = append(container, fmt.Sprintf("%s %s", coffee, "ready :)"))
	}

	return
}

func barista(order string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", order, "espresso")
}
