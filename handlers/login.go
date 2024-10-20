package handlers

import (
    "html/template"
    "net/http"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Login(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    if r.Method == "POST" {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")
        // Handle login logic here
        if username == "admin" && password == "admin" {
            session.Values["authenticated"] = true
            session.Save(r, w)
            http.Redirect(w, r, "/info", http.StatusFound)
            return
        }
    }
    tmpl := template.Must(template.ParseFiles("templates/login.html"))
    tmpl.Execute(w, nil)
}

