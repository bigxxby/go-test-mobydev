package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func (mHandler Main_handler) AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie, err := r.Cookie("sessionID")
		sessionId := cookie.Value
		user, authorized, err := mHandler.Data.IsAuthorised(sessionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !authorized || user.IsAdmin.Int16 != 1 {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
		} else {

			temp, err := template.ParseFiles("ui/templates/admin.html")
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
}
