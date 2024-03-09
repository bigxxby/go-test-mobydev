package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var userData struct {
		UserID  string `json:"user_id"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		DOB     string `json:"date_of_birth"`
		IsAdmin bool   `json:"is_admin"`
	}

	if err := json.Unmarshal(body, &userData); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	err = mHandler.Data.UpdateUserAdmin(userData.UserID, userData.Name, userData.Phone, userData.DOB, userData.IsAdmin)
	if err != nil {
		http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("User with id:", userData.UserID, "updated")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User data updated successfully"))
}

func (mHandler Main_handler) DeleteUserAdminHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	err := mHandler.Data.DeleteUser(userID)
	if err != nil {
		http.Error(w, "Error deleting user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
