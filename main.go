package main

import (
	"goTweet/routers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routers.UserRouter(r)
	routers.TweetRouter(r)

	http.ListenAndServe(":8080", r)
}
