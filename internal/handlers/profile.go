package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func (mHandler Main_handler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("sessionID")
	sessionId := cookie.Value

	user, authorized, err := mHandler.Data.IsAuthorised(sessionId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !authorized {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
	} else {
		temp, err := template.ParseFiles("ui/templates/profile.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = temp.Execute(w, user)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}
