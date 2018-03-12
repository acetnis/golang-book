package home

import (
	"encoding/json"
	"net/http"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

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
