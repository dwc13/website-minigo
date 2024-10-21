package handlers

import (
    "net/http"
    "website-minigo/database"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        id := r.FormValue("id")
        newUsername := r.FormValue("newusername")

        var user database.User
        result := database.DB.First(&user, id)
        if result.Error == nil {
            user.Username = newUsername
            database.DB.Save(&user)
        }
        http.Redirect(w, r, "/admin", http.StatusFound)
    }
}

