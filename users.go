
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UsersHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		fmt.Fprintln(response, "Method not supported")
		return
	}
	// user credentials
	user := []User{
		{
			Email: 	  "admin@example.com",
			Password: "admin",
		},
		{
			Email: 	  "admin1@example.com",
			Password: "admin1",
		},
	}
	out, err := json.Marshal(user)
	if err != nil {
		Log.Println(err.Error())
		return
	}
	fmt.Fprintln(response, string(out))
}
