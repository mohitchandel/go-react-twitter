package controller

import (
	"encoding/json"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{user: user}")
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{user: user}")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{user: user}")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{user: user}")
}
