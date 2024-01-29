package main

import (
	"exam2game/internal/handlers"
	"exam2game/pkg/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dbHost := "localhost"
	dbName := "postgres"
	dbUser := "rahul"
	dbPort := "5432"
	dbPassword := "12345"

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", dbHost, dbUser, dbName, dbPort, dbPassword)

	db := database.InitDB(connStr)
	defer db.Close()

	router := mux.NewRouter()

	handlers.RegisterCourseHandlers(router, db)

	log.Fatal(http.ListenAndServe(":8080", router))
}
