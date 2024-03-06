package main

import (
	"fmt"
	gotest "gotest/internal"
	"log"
	"net/http"
)

func main() {
	database, err := gotest.NewDatabase("database.db")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer database.Close()
	err = database.CreateProjectsTable()
	err = database.CreateUsersTable()
	if err != nil {
		log.Println("Error creating Table :", err.Error)
		return
	}

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", gotest.IndexHandler)
	http.HandleFunc("/reg", gotest.RegHandler)
	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
