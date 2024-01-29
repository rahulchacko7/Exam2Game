package handlers

import (
	"database/sql"
	"encoding/json"
	"exam2game/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterCourseHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/courses", getCourses(db)).Methods("GET")
	router.HandleFunc("/courses/{id}", getCourse(db)).Methods("GET")
	router.HandleFunc("/courses", createCourse(db)).Methods("POST")
	router.HandleFunc("/courses/{id}", updateCourse(db)).Methods("PUT")
	router.HandleFunc("/courses/{id}", deleteCourse(db)).Methods("DELETE")
}

func getCourses(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var courses []models.Course
		rows, err := db.Query("SELECT id, title FROM courses")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var course models.Course
			if err := rows.Scan(&course.ID, &course.Title); err != nil {
				log.Fatal(err)
			}
			courses = append(courses, course)
		}

		json.NewEncoder(w).Encode(courses)
	}
}

func getCourse(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var course models.Course
		err := db.QueryRow("SELECT id, title FROM courses WHERE id = $1", params["id"]).Scan(&course.ID, &course.Title)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(course)
	}
}

func createCourse(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var course models.Course
		json.NewDecoder(r.Body).Decode(&course)
		_, err := db.Exec("INSERT INTO courses(title) VALUES($1)", course.Title)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func updateCourse(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var course models.Course
		json.NewDecoder(r.Body).Decode(&course)
		_, err := db.Exec("UPDATE courses SET title = $1 WHERE id = $2", course.Title, params["id"])
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteCourse(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		_, err := db.Exec("DELETE FROM courses WHERE id = $1", params["id"])
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
