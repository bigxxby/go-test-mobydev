package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func (mHandler Main_handler) ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	temp, err := template.ParseFiles("ui/templates/projects.html")
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

func (mHandler Main_handler) CreateProjectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie, err := r.Cookie("sessionID")
		sessionId := cookie.Value
		user, authorized, err := mHandler.Data.IsAuthorised(sessionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !authorized {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}

		temp, err := template.ParseFiles("ui/templates/createProject.html")
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = temp.Execute(w, user.Id)
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		userIdStr := r.Form.Get("userId")
		yearStr := r.Form.Get("year")
		durationStr := r.Form.Get("duration")
		name := r.Form.Get("name")
		category := r.Form.Get("category")
		projectType := r.Form.Get("projectType")
		ageCategory := r.Form.Get("ageCategory")
		keywords := r.Form.Get("keywords")
		description := r.Form.Get("description")
		director := r.Form.Get("director")
		producer := r.Form.Get("producer")

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			log.Println("Error converting userId to integer:", err)
			http.Error(w, "Invalid userId", http.StatusBadRequest)
			return
		}
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			log.Println("Error converting year to integer:", err)
			http.Error(w, "Invalid year", http.StatusBadRequest)
			return
		}
		duration, err := strconv.Atoi(durationStr)
		if err != nil {
			log.Println("Error converting duration to integer:", err)
			http.Error(w, "Invalid duration", http.StatusBadRequest)
			return
		}

		err = mHandler.Data.AddProject(userId, year, duration, name, category, projectType, ageCategory, keywords, description, director, producer)
		if err != nil {
			log.Println("Error adding project:", err)
			http.Error(w, "Failed to add project", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Project added successfully")
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
}

func (mHandler Main_handler) GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := mHandler.Data.GetAllProjects()
	if err != nil {
		http.Error(w, "Failed to retrieve projects", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(projects)
	if err != nil {
		http.Error(w, "Failed to encode projects to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
