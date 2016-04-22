
package main

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	},
	// debug mode activated!
	Debug: false,
})

func main() {
	router := http.NewServeMux()

	// links `api/admin` route to protected handler
	router.HandleFunc("/api/login", AuthHandler)
	router.Handle("/api/users", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(UsersHandler)),
	))

	router.Handle("/api/authenticated", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(AuthenticateHandler)),
	))
	
	// serves the static files for angular js application as well as all other resources
	router.Handle("/", http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8080", router)
}
