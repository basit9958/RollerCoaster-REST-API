package main

import (
	"encoding/json"
	"net/http"
)

type Coaster struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID           string `json:"id"`
	InPark       string `json:"inPark"`
	Height       string `json:"height"`
}

type coasterHandlers struct {
	store map[string]Coaster
}


func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))
	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1": Coaster{
				Name:         "Fury 311",
				Manufacturer: "BMW",
				ID:           "id1",
				InPark:       "carowinds",
				Height:       "15 feet",
			},
		},
	}
}
func main() {
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.get)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}
