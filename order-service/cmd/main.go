package main

import (
	"log"
	"net/http"
	"Order-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/order/notify", handlers.OrderNotify).Methods("POST")

	log.Println("Order Service is running on port 9001")
	log.Fatal(http.ListenAndServe(":9001", r))
}
