package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"time"
)

func Sleep(w http.ResponseWriter, r *http.Request) {
	d := time.Millisecond * time.Duration(rand.Int64N(1000))
	time.Sleep(d)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	type ResponseSchema struct {
		Time string `json:"timeTaken"`
	}
	json.NewEncoder(w).Encode(ResponseSchema{Time: d.String()})
}
