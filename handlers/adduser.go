package handlers

import (
    "net/http"
    "website-minigo/database"
    "golang.org/x/crypto/bcrypt"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")
        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        user := database.User{Username: username, Password: string(hashedPassword)}
        database.DB.Create(&user)
        http.Redirect(w, r, "/admin", http.StatusFound)
    }
}

