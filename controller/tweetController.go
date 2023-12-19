package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Tweet struct {
	Title     string
	Body      string
	Author    uint64
	CreatedAt string
}

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")
	author := r.URL.Query().Get("author")
	authorId, err := strconv.ParseUint(author, 10, 32)
	createdAt := r.URL.Query().Get("created_at")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tweet := Tweet{
		Title:     title,
		Body:      body,
		Author:    authorId,
		CreatedAt: createdAt,
	}
	json.NewEncoder(w).Encode(tweet)
}

func EditTweet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{user: user}")
}

func GetTweets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{user: user}")
}

func GetTweet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{user: user}")
}
