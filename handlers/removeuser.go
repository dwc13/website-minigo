package handlers

import (
    "net/http"
    "website-minigo/database"
)

func RemoveUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        id := r.FormValue("id")
        username := r.FormValue("username")
        database.DB.Where("id = ? AND username = ?", id, username).Delete(&database.User{})
        http.Redirect(w, r, "/admin", http.StatusFound)
    }
}

