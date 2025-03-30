package handler

import (
	"encoding/json"
	"fmt"
	"gocloud/internals/models"
	"gocloud/internals/services"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
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
func UploadFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	req.ParseMultipartForm(10 << 20)
	file, handlr, err := req.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handlr.Filename)
	fmt.Printf("File Size: %+v\n", handlr.Size)
	fmt.Printf("MIME Header: %+v\n", handlr.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

var conn = services.GetDb()
