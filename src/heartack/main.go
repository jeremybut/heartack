package main

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"

	"heartack/controllers"

	"database/sql"
	"fmt"
	"net/http"
	"time"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	},
	// debug mode activated!
	Debug: false,
})

func initDB() {
	fmt.Println("Initializing database")
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrate",
	}

	db, err := sql.Open("postgres", "dbname=heartack sslmode=disable")
	if err != nil {
		// Handle errors!
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func main() {
	initDB()

	fmt.Println("=> Booting GoServer")
	fmt.Println("=> Go application starting on http://0.0.0.0:3000")
	fmt.Println("=> Ctrl-C to shutdown server")
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

	// serves the static files for angular js application as well as all other resources
	router.Handle("/", http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8080", router)
}