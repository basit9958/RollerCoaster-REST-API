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

func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		h.get(w,r)
		return
	case "POST":
		h.post(w,r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

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
func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request){
	
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
	http.HandleFunc("/coasters", coasterHandlers.coasters)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}
