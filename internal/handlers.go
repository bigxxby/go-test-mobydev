package gotest

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/mail"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("ui/templates/index.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = temp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/login" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	if r.Method == "POST" {
// 		return
// 	} else if r.Method == "GET" {
// 		temp, err := template.ParseFiles("ui/templates/login.html")
// 		if err != nil {
// 			log.Println(err.Error())
// 			return
// 		}
// 		err = temp.Execute(w, nil)
// 		if err != nil {
// 			log.Println(err.Error())
// 			return
// 		}
// 		return
// 	}
// }
func RegHandler(w http.ResponseWriter, r *http.Request) {
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
		isValid := IsValidPassword(password) // if not valid password
		if !isValid {
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "User %s successfully registered!", email)
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
	// ADD USER HERE
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int) {
	w.Header().Set("Content-Type", "application/json")

	response := fmt.Sprintf(`{"error": "Error %d"}`, code)

	w.WriteHeader(code)
	_, err := w.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
