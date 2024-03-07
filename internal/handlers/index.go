package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func (mHandler Main_handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
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
