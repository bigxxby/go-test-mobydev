package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func (mHandler Main_handler) LogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		if email == "" || password == "" {
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		user, exists, err := mHandler.Data.CheckUser(email, password)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("Error checking user:", err)
			return
		}

		if !exists {
			log.Println("User does not exist or password is incorrect")
			http.Error(w, "User not found or password incorrect", http.StatusUnauthorized)
			return
		}

		cookie := http.Cookie{
			Name:     "sessionID",
			Value:    user.SessionId.String,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   86400, // 1 day expiration
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	} else if r.Method == "GET" {
		temp, err := template.ParseFiles("ui/templates/login.html")
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("Error parsing login template:", err)
			return
		}
		err = temp.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("Error executing login template:", err)
			return
		}
		return

	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
}
