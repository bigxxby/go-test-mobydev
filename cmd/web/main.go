package main

import (
	// gotest "gotest/internal"

	"gotest/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// database, err := gotest.NewDatabase("database.db")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer database.Close()
	// err = database.CreateProjectsTable()
	// err = database.CreateUsersTable()
	// if err != nil {
	// 	log.Println("Error creating Table :", err.Error)
	// 	return
	// }
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

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", mHandler.IndexHandler)
	http.HandleFunc("/reg", mHandler.RegHandler)
	http.HandleFunc("/login", mHandler.LogHandler)
	http.HandleFunc("/profile", mHandler.ProfileHandler)
	http.HandleFunc("/logout", mHandler.LogoutHandler)
	// http.HandleFunc("/projects", mHandler.ProjectsHandler)
	http.HandleFunc("/createProject", mHandler.CreateProjectsHandler)
	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
