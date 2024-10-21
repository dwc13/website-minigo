package handlers

import (
    "html/template"
    "net/http"
    "website-minigo/database"
)

type AdminPageData struct {
    Users []database.User
}

func Admin(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    if isAdmin, ok := session.Values["isAdmin"].(bool); !ok || !isAdmin {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    var users []database.User
    database.DB.Find(&users)

    tmpl := template.Must(template.ParseFiles("templates/admin.html"))
    data := AdminPageData{Users: users}
    tmpl.Execute(w, data)
}

