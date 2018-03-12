package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "helle, World!")
		})

	barHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Bar!")
	}
	http.HandleFunc("/bar", barHandler) //http://localhost:3000/bar

	http.ListenAndServe(":3000", nil) //http://localhost:3000
}
