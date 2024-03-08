package handlers

import (
	"encoding/json"
	"net/http"
)

func (mHandler Main_handler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := mHandler.Data.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to get all users", http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to marshal users to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(usersJSON)
}

func (mHandler Main_handler) UpdateUserAdminHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	userID := r.Form.Get("userId")
	newName := r.Form.Get("name")
	newPhone := r.Form.Get("phone")
	newDOB := r.Form.Get("date_of_birth")
	isAdmin := r.Form.Get("is_admin") == "on"

	err = mHandler.Data.UpdateUserAdmin(userID, newName, newPhone, newDOB, isAdmin)
	if err != nil {
		http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
