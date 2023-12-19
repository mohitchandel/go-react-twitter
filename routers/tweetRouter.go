package routers

import (
	"goTweet/controller"

	"github.com/gorilla/mux"
)

func TweetRouter(router *mux.Router) {
	router.HandleFunc("/api/tweet/create", controller.CreateTweet).Methods("POST")
	router.HandleFunc("/api/tweets", controller.GetTweets).Methods("GET")
	router.HandleFunc("/api/tweet/:id", controller.EditTweet).Methods("PUT")
	router.HandleFunc("/api/tweet/:id", controller.GetTweet).Methods("GET")
}
