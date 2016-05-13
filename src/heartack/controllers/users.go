package controllers

import (
	"heartack/models"
	"heartack/utils"

	"encoding/json"
	"fmt"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	pts, err := models.AllPatients()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	out, err := json.Marshal(pts)
	if err != nil {
		utils.Log.Println(err.Error())
		return
	}
	fmt.Fprintln(w, string(out))
}
