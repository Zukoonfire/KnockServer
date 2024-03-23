package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

type Response struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	port := ":8080" // Change the port if necessary
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if msg.Message == "Knock, Knock" {
		response := Response{Response: "Who's there?"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response := Response{Error: "Invalid message"}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
