package handlers

import (
	"gotest/internal/logic"
	"log"
	"net/http"
	"net/mail"
	"text/template"
)

func (mHandler Main_handler) RegHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reg" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")
		if email == "" || confirmPassword == "" || password == "" { // if empty
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}
		if password != confirmPassword { // if not match
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}
		_, err := mail.ParseAddress(email) // if not valid email
		if err != nil {
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}
		isValid := logic.IsValidPassword(password) // if not valid password
		if !isValid {
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		err = mHandler.Data.AddUser(email, password)
		if err != nil {

			log.Println("Error adding user: ", err.Error())
			if err.Error() == "UNIQUE constraint failed: users.email" {
				_, err := w.Write([]byte("User already exists"))
				if err != nil {
					http.Error(w, "Error writing response", http.StatusInternalServerError)
					return
				}
				return
			}
			return
		}
		log.Println("User %s successfully registered!", email)
		http.Redirect(w, r, "/login", http.StatusSeeOther)

	} else if r.Method == "GET" {
		temp, err := template.ParseFiles("ui/templates/reg.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
}
