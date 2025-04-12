package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ReqData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ResData struct {
	Message string `json:"message"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	var data ReqData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := ResData{Message: fmt.Sprintf("Hello %s, age %d!", data.Name, data.Age)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/submit", handle)
	log.Println("🚀 Server running on http://localhost:10102")
	log.Fatal(http.ListenAndServe(":10102", nil))
}
