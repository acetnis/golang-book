package main

import (
	"fmt"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

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

	http.Handle("/home", &HomePageHandler{}) //http://localhost:3000/home

	http.ListenAndServe(":3000", nil) //http://localhost:3000
}
