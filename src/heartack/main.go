package main

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"heartack/controllers"
	"heartack/models"
 
  "encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func main() {
	models.InitDB("postgres://jeremy@localhost/heartack?sslmode=disable")
	fmt.Println("Server Started at: ", time.Now())

	router := http.NewServeMux()

	// links `api/admin` route to protected handler
	router.HandleFunc("/api/login", controllers.AuthHandler)
	router.Handle("/api/users", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(controllers.UsersHandler)),
	))

	router.Handle("/api/authenticated", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(controllers.AuthenticateHandler)),
	))

	serves the static files for angular js application as well as all other resources
	router.Handle("/", http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8080", router)
}
