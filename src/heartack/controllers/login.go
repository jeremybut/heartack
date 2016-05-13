package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"heartack/models"
	"heartack/utils"
	"net/http"
	"time"
)

// specific handle for authing the API for the website
func AuthHandler(resp http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		fmt.Fprintln(resp, "Method not supported")
		return
	}

	decoder := json.NewDecoder(request.Body)
	cred := models.User{"", "", "", ""}
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

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "All good. You only get this message if you're authenticated")
}
