package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "website-minigo/database"
    "website-minigo/handlers"
)

func main() {
    database.InitDatabase()
    database.Seed()  // Seed the database with initial users

    r := mux.NewRouter()
    r.HandleFunc("/", handlers.Login).Methods("GET", "POST")
    r.HandleFunc("/info", handlers.Info).Methods("GET")
    r.HandleFunc("/logout", handlers.Logout).Methods("POST")
    r.HandleFunc("/adduser", handlers.AddUser).Methods("POST")
    r.HandleFunc("/removeuser", handlers.RemoveUser).Methods("POST")
    r.HandleFunc("/admin", handlers.Admin).Methods("GET")
    r.HandleFunc("/welcome", handlers.Welcome).Methods("GET")
    r.HandleFunc("/edituser", handlers.EditUser).Methods("POST")
    http.Handle("/", r)

    http.ListenAndServe(":8080", r)
}

