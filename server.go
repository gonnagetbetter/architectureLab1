package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"time"`
}

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		currentTime := time.Now().Format(time.RFC3339)
		response := TimeResponse{currentTime}

		jsonBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})

	err := http.ListenAndServe(":8795", nil)
	if err != nil {
		panic(err)
	}
}
