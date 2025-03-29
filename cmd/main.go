package main

import (
	"gocloud/internals/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/registor", handler.RegisterUser)
	log.Println("Server runnning on PORT", 8085)
	err := http.ListenAndServe("localhost:8085", nil)
	if err != nil {
		log.Println("Error in starting the server", err)
		return
	}
}
