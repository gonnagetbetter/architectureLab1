package main

import (
	"encoding/json"
	"net/http"
	"time"
)
  
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
  }

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
}
