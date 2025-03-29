package main

import (
	"gocloud/internals/handler"
	"log"
	"net/http"
	"gocloud/internals/services"
	"fmt"
	"gocloud/internals/models"
)

func main() {
	http.HandleFunc("/registor", handler.RegisterUser)
	log.Println("Server runnning on PORT", 8085)
	err := http.ListenAndServe("localhost:8085", nil)
	if err != nil {
		log.Println("Error in starting the server", err)
		return
	}

	conn := services.GetDb()

	err := services.InsertUser(conn, models.User{Name: "om", Email: "agraw333a55l", Password: "om"})
	if err != nil {
		fmt.Println(err)
		return
	}
	email, err := services.GetUserByEmail(conn, "agrawahhgygul")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(email)
}