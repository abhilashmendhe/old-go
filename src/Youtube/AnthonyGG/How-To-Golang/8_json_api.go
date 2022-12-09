package main

import (
	"encoding/json"
	"net/http"
)

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {

	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	w.Header().Add("Content-Type", "application/json")
	// 	w.Write([]byte("not allowed"))
	// 	return
	// }
	if r.Method != http.MethodGet{
		if err := writeJSON(w,http.StatusMethodNotAllowed, "not allowed")
		return 
	}

}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func main() {

	http.HandleFunc("/user", handleGetUserByID)
	http.ListenAndServe(":3000", nil)
}
