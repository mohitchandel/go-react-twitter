package controller

import (
	"encoding/json"
	"fmt"
	"goTweet/database"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint64
	UserName  string
	Email     uint64
	Password  string
	CreatedAt string
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	hashedPassword, err := hashPassword(password)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

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
	insertQuery := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	_, err = db.Exec(insertQuery, username, email, hashedPassword)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprintf(w, "user created successfully")
	json.NewEncoder(w).Encode("{user: user}")
}

var store = sessions.NewCookieStore([]byte("ss11"))

func LoginUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var storedPassword string
	// Query the database to retrieve the stored password for the given username
	err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// If the passwords match, create a session for the user
	session, _ := store.Get(r, "session-name")
	session.Values["username"] = username
	session.Save(r, w)

	fmt.Fprintf(w, "User logged in successfully")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	selectQuery := "SELECT id, username, email, created_at FROM users WHERE id = ?"
	rows, err := db.Query(selectQuery, userIDInt)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var user User

	if rows.Next() {
		if err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.CreatedAt); err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to scan user data", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	selectQuery := "SELECT * FROM users"
	rows, err := db.Query(selectQuery)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.CreatedAt); err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to scan user data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error during row iteration", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
