package handlers

import (
    "net/http"
    "website-minigo/database"
)

func RemoveUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        username := r.FormValue("username")
        database.DB.Where("username = ?", username).Delete(&database.User{})
        http.Redirect(w, r, "/info", http.StatusFound)
    }
}

