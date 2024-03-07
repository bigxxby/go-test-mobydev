package handlers

import (
	"log"
	"net/http"
)

func (mHandler Main_handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, 500)
		return
	}
	err = mHandler.Data.LogoutUser(cookie.Value)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, 500)
		return
	}
}
