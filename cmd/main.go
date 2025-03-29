package main

import (
	"fmt"
	"gocloud/internals/models"
	"gocloud/internals/services"
)

func main() {
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
