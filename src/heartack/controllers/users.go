package controllers

import (
	"encoding/json"
	"fmt"
	"heartack/models"
	"net/http"
	"heartack/utils"
)

func UsersHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		fmt.Fprintln(response, "Method not supported")
		return
	}
	// user credentials
	user := []models.User{
		{
			Email:    "admin@example.com",
			Password: "admin",
		},
		{
			Email:    "admin1@example.com",
			Password: "admin1",
		},
	}
	out, err := json.Marshal(user)
	if err != nil {
		utils.Log.Println(err.Error())
		return
	}
	fmt.Fprintln(response, string(out))
}
