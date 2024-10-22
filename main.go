package main

import (
	"os"
	"os/signal"
	"syscall"
    "net/http"
    "github.com/gorilla/mux"
    "website-minigo/database"
    "website-minigo/handlers"
)

func main() {
    database.InitDatabase()

	// Set up a channel to listen for interrupt signals
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    // Run cleanup on signal interrupt
    go func() {
        <-c
        database.CloseDatabase()
        os.Exit(0)
    }()

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

