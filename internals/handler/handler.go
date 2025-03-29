package handler

import (
	"encoding/json"
	"gocloud/internals/models"
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
	w.Write([]byte("Hello World"))
	if req.Method == "POST" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err2 := w.Write([]byte("Request body not supported"))
			if err2 != nil {
				return
			}
		}
		re := models.NewUser()
		err = json.Unmarshal(body, &re)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		hashedpass, _ := HashPassword(re.Password)
		re.Password = hashedpass
		log.Println(re)
	}
}
func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))

}
