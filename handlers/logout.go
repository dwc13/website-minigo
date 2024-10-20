package handlers

import (
    "net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    session.Values["authenticated"] = false
    session.Save(r, w)
    http.Redirect(w, r, "/", http.StatusFound)
}

