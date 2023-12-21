package routers

import (
	"goTweet/controller"

	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router) {
	router.HandleFunc("/api/user/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/api/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/", controller.GetUser).Methods("GET")
	router.HandleFunc("/api/user/login", controller.LoginUser).Methods("POST")
}
