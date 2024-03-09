package main

import (
	// gotest "gotest/internal"

	"gotest/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mHandler := handlers.Init_handler()
	if mHandler == nil {
		return
	}
	err := mHandler.Data.CreateProjectsTable()
	err = mHandler.Data.CreateUsersTable()
	if err != nil {
		log.Println("Error creating tables : ", err.Error())
		return
	}
	/// ОСТАНОВИЛСЯ НА Uncaught (in promise) ReferenceError: response is not defined
	// <anonymous> http://localhost:8080/admin:101
	// promise callback* http://localhost:8080/admin:98
	// EventListener.handleEvent* http://localhost:8080/admin:70
	// EventListener.handleEvent* http://localhost:8080/admin:59
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", mHandler.IndexHandler)
	http.HandleFunc("/reg", mHandler.RegHandler)
	http.HandleFunc("/login", mHandler.LogHandler)
	http.HandleFunc("/profile", mHandler.ProfileHandler)
	http.HandleFunc("/logout", mHandler.LogoutHandler)
	http.HandleFunc("/projects", mHandler.ProjectsHandler)
	http.HandleFunc("/createProject", mHandler.CreateProjectsHandler)
	http.HandleFunc("/getProjects", mHandler.GetProjectsHandler)
	http.HandleFunc("/updateProfile", mHandler.UpdateProfileHandler)
	http.HandleFunc("/admin", mHandler.AdminHandler)
	http.HandleFunc("/getAllUsers", mHandler.GetAllUsersHandler)
	http.HandleFunc("/updateUserAdmin", mHandler.UpdateUserAdminHandler)
	http.HandleFunc("/deleteUser/", mHandler.DeleteUserAdminHandler)
	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
