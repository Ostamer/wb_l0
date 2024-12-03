package main

import (
	"log"
	"net/http"
)

func main() {

	initConfig()
	initDB()
	initCache()
	go startKafkaConsumer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/order", getOrderHandler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
