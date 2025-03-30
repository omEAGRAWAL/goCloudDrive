package handler

import (
	"encoding/json"
	"fmt"
	"gocloud/internals/models"
	"gocloud/internals/services"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("On register "))
	if req.Method == "POST" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err2 := w.Write([]byte("Request body not supported"))
			if err2 != nil {
				return
			}
		}
		usermod := models.NewUser()
		err = json.Unmarshal(body, &usermod)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		hashedpass, _ := HashPassword(usermod.Password)
		usermod.Password = hashedpass
		log.Println(usermod)
		err2 := services.InsertUser(conn, *usermod)
		fmt.Println("stored in db")
		if err2 != nil {
			return
		}

	}
}
func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("on login page"))
	if req.Method == "POST" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		//services.GetUserByEmail(conn,bo)
		re := models.NewUser()
		err = json.Unmarshal(body, &re)
		fmt.Println(re)
		data, err := services.GetUserByEmail(conn, re.Email)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)

	}

}

var conn = services.GetDb()
