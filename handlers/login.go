package handlers

import (
    "html/template"
    "net/http"
    "website-minigo/database"
    "golang.org/x/crypto/bcrypt"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Login(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    var errorMessage string
    if r.Method == "POST" {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")

        var user database.User
        result := database.DB.Where("username = ?", username).First(&user)
        if result.Error != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
            errorMessage = "Invalid username or password"
        } else {
            session.Values["authenticated"] = true
            session.Values["isAdmin"] = user.IsAdmin
            session.Save(r, w)

            if user.IsAdmin {
                http.Redirect(w, r, "/admin", http.StatusFound)
            } else {
                http.Redirect(w, r, "/welcome", http.StatusFound)
            }
            return
        }
    }
    tmpl := template.Must(template.ParseFiles("templates/login.html"))
    tmpl.Execute(w, errorMessage)
}

