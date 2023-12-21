# Twitter-like Backend Project in Golang

This project is a backend implementation resembling a Twitter-like platform built using Golang. It enables user registration, login, tweet creation, retrieval of all tweets, and deletion of tweets.

## Features

- User Registration: Users can sign up for an account.
- User Authentication: Registered users can log in securely.
- Tweet Creation: Authenticated users can create new tweets.
- Get All Tweets: Retrieve all tweets available.
- Delete Tweets: Authenticated users can delete their own tweets.

## Technologies Used

- Golang
- MySQL

## Installation

1. Clone the repository: `git clone https://github.com/mohitchandel/go-twitter.git`
2. Navigate to the project directory.
3. Install dependencies: `go mod download`

## Usage

1. Set up a MySQL database and configure the connection in a configuration file.
2. Create a configuration file or environment variables for sensitive information.
3. Start the server: `go run main.go` or compile and execute the binary.

## API Endpoints

### Authentication

- `POST /api/user/register`: Register a new user.
- `POST /api/user/login`: Log in as a registered user.
- `GET /api/users`: Get all users.
- `GET /api/user`: Get a single user (param user_id).

### Tweets

- `GET /api/tweets`: Get all tweets.
- `POST /api/tweet/create`: Create a new tweet.
- `DELETE /api/tweet`: Delete a specific tweet by ID.
- `GET /api/tweet`: Get a single tweet (param tweet_id).
