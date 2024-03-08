package handlers

import (
	"encoding/json"
	"gotest/internal/database"
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

func (mHandler Main_handler) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, "Session cookie not found", http.StatusBadRequest)
		return
	}
	sessionID := cookie.Value

	_, authorized, err := mHandler.Data.IsAuthorised(sessionID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !authorized {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}

	var updateReq database.UpdateUser
	err = json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		log.Println(err.Error())

		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	err = mHandler.Data.UpdateUser(updateReq.Id, updateReq.Name, updateReq.Phone, updateReq.DateOfBirth)
	if err != nil {
		log.Println(err.Error())

		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
