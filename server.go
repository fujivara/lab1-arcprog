package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	response := TimeResponse{Time: currentTime}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", fooHandler)
	// go go go go
	// go go go go
	log.Fatal(http.ListenAndServe(":9875", nil))
}
