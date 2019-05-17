package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const url = "http://localhost:8081"

type Apple struct {
	ID    int    `json:"id"`
	Fresh bool   `json:"fresh"`
	Color string `json:"color"`
}

func appleHandler1(w http.ResponseWriter, r *http.Request) {
	log.Printf("apple handler 1 called")
	apple := Apple{
		ID:    1,
		Fresh: true,
		Color: "green",
	}
	b, _ := json.Marshal(apple)
	w.Write(b)
}

func appleHandler2(w http.ResponseWriter, r *http.Request) {
	log.Printf("apple handler 2 called")
	apple := Apple{
		ID:    2,
		Fresh: true,
		Color: "red",
	}
	b, _ := json.Marshal(apple)
	w.Write(b)
}

func applesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("apples called")
	apples := []Apple{
		{
			ID:    1,
			Fresh: true,
			Color: "green",
		},
		{
			ID:    2,
			Fresh: false,
			Color: "red",
		},
	}
	b, _ := json.Marshal(apples)
	w.Write(b)
}

func main() {
	http.HandleFunc("/apple/1", appleHandler1)
	http.HandleFunc("/apple/2", appleHandler2)
	http.HandleFunc("/apples", applesHandler)
	http.ListenAndServe(":8081", nil)
}
