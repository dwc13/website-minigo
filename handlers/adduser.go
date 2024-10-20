package handlers

import (
    "net/http"
    "website-minigo/database"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        user := database.User{Username: r.FormValue("username"), Password: r.FormValue("password")}
        database.DB.Create(&user)
        http.Redirect(w, r, "/info", http.StatusFound)
    }
}

