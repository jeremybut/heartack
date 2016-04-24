package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"heartack/models"
	"net/http"
	"time"
	"heartack/utils"
)

// specific handle for authing the API for the website
func AuthHandler(resp http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		fmt.Fprintln(resp, "Method not supported")
		return
	}
	decoder := json.NewDecoder(request.Body)
	cred := models.User{"test", "test", "ok", "ok"}
	err := decoder.Decode(&cred)

	if err != nil {
		utils.Log.Println(err.Error())
		return
	}

	if cred.Email == "admin@example.com" && cred.Password == "admin" {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims["foo"] = "bar"
		token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			utils.Log.Println(err.Error())
		}
		resp.Header().Add("X-Auth-Token", tokenString)
	} else {
		fmt.Fprintln(resp, "This is an authenticated request")
	}
}

// specific handle for accessing the admin API for the website
func AuthenticateHandler(resp http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		fmt.Println(resp, "Method not supported")
		return
	}
	fmt.Println(resp, "true")
}
