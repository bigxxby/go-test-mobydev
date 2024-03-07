package handlers

import (
	"gotest/internal/database"
	"log"
)

func Init_handler() *Main_handler {
	mHandler := Main_handler{}
	var err error
	mHandler.Data, err = database.Init_DB()
	if err != nil {
		log.Println("Error creating Struct :", err.Error())
		return nil
	}
	return &mHandler
}
