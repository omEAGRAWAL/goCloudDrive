package main

import (
	"gocloud/internals/handler"
	"log"
	"net/http"
)

func main() {

	//conn := services.GetDb()

	//err := services.InsertUser(conn, models.User{Name: "om", Email: "a5555l", Password: "om"})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//email, err := services.GetUserByEmail(conn, "agraw333a55555l")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(email)
	fs := http.FileServer(http.Dir("internals/client/gocloudui/dist"))
	// Handle root by serving index.html explicitly
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/upload", handler.UploadFile)
	http.HandleFunc("/api/register", handler.RegisterUser)
	http.HandleFunc("/api/login", handler.Login)
	log.Println("Server runnning on PORT", 8085)
	err1 := http.ListenAndServe("localhost:8085", nil)
	if err1 != nil {
		log.Println("Error in starting the server", err1)
		return
	}

}
