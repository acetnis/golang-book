package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "addlication/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(user)
	w.Write(data)
}

func main() {
	http.Handle("/", &HomePageHandler{}) //http://localhost:3000/?name=xxx

	http.ListenAndServe(":3000", nil) //http://localhost:3000
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}
