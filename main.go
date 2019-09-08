package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response struct {
	Message string
	Status  int64
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	message := Response{
		Message: "Welcome ❤️",
		Status:  200,
	}

	_ = json.NewEncoder(w).Encode(message)
}

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":9090", route))
}
