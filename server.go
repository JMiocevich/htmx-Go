package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/graphdata", getGraphData).Methods("GET")
	r.HandleFunc("/", serveHomePage).Methods("GET")
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Handle CORS
	handler := cors.Default().Handler(r)

	http.ListenAndServe(":8080", handler)
}

func getGraphData(w http.ResponseWriter, r *http.Request) {
	data := map[string][]int{
		"data": {5, 10, 15, 20},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
