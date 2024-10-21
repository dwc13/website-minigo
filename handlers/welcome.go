package handlers

import (
    "html/template"
    "net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    tmpl := template.Must(template.ParseFiles("templates/welcome.html"))
    tmpl.Execute(w, nil)
}

