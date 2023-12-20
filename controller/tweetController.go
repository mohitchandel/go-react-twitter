package controller

import (
	"encoding/json"
	"fmt"
	"goTweet/database"
	"net/http"
	"strconv"
)

type Tweet struct {
	Id        string
	Title     string
	Body      string
	Author    uint64
	CreatedAt string
}

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")
	author := r.URL.Query().Get("author_id")
	authorId, err := strconv.ParseUint(author, 10, 64)

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Get the database connection
	insertQuery := "INSERT INTO tweets (title, body, author_id) VALUES (?, ?, ?)"
	_, err = db.Exec(insertQuery, title, body, authorId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprintf(w, "user created successfully")
}

func GetTweets(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	selectQuery := "SELECT * FROM tweets"
	rows, err := db.Query(selectQuery)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to fetch tweets", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tweets []Tweet

	for rows.Next() {
		var tweet Tweet
		if err := rows.Scan(&tweet.Id, &tweet.Title, &tweet.Body, &tweet.Author, &tweet.CreatedAt); err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to scan tweet data", http.StatusInternalServerError)
			return
		}
		tweets = append(tweets, tweet)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error during row iteration", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tweets)
}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("tweet_id")

	// Convert the tweetID to an integer for safety and use it in the query
	tweetIDInt, err := strconv.Atoi(tweetID)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	deleteQuery := "DELETE FROM tweets WHERE id = ?"
	_, err = db.Exec(deleteQuery, tweetIDInt)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to delete tweet", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Tweet with ID %s deleted successfully", tweetID)
}

func GetTweet(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("tweet_id")
	tweetIDInt, err := strconv.Atoi(tweetID)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	selectQuery := "SELECT id, title, body, author, created_at FROM tweets WHERE id = ?"
	rows, err := db.Query(selectQuery, tweetIDInt)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to fetch tweet", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tweet Tweet

	if rows.Next() {
		if err := rows.Scan(&tweet.Id, &tweet.Title, &tweet.Body, &tweet.Author, &tweet.CreatedAt); err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to scan tweet data", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Tweet not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tweet)
}
