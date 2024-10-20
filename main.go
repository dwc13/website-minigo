package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "website-minigo/handlers"
	"website-minigo/database"
)

func main() {
	database.InitDatabase()
    r := mux.NewRouter()
    r.HandleFunc("/", handlers.Login).Methods("GET", "POST")
    r.HandleFunc("/info", handlers.Info).Methods("GET")
    r.HandleFunc("/logout", handlers.Logout).Methods("POST")
    r.HandleFunc("/adduser", handlers.AddUser).Methods("POST")
    r.HandleFunc("/removeuser", handlers.RemoveUser).Methods("POST")
    http.Handle("/", r)

    http.ListenAndServe(":8088", r)
}

