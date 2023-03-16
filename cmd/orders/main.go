package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Benjamin-G/Asynchronous-Event-Handling/src/orders"
)

func main() {
	fmt.Println("Orders Server Started...")
	fmt.Println("Listening on http://localhost:8080")
	http.HandleFunc("/healthcheck", orders.HealthCheck)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
