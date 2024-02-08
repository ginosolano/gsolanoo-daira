package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Entry struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Operacion string `json:"operation"`
	Resultado string `json:"result"`
}

var history = []Entry{}

func main() {
	http.HandleFunc("/api/history", historyHandler)
	http.Handle("/", http.FileServer(http.Dir("public"))) 

	fmt.Println("Server is running on :5173...")
	http.ListenAndServe(":5173", nil)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	if r.Method == http.MethodGet {
		
		json.NewEncoder(w).Encode(history)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}