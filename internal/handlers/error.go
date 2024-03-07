package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int) {
	w.Header().Set("Content-Type", "application/json")

	response := fmt.Sprintf(`{"error": "Error %d"}`, code)

	w.WriteHeader(code)
	_, err := w.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
