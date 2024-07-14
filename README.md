# GlobalWebIndex Engineering Challenge

## Introduction

This challenge is designed to give you the opportunity to demonstrate your abilities as a software engineer and specifically your knowledge of the Go language.

On the surface the challenge is trivial to solve, however you should choose to add features or capabilities which you feel demonstrate your skills and knowledge the best. For example, you could choose to optimise for performance and concurrency, you could choose to add a robust security layer or ensure your application is highly available. Or all of these.

Of course, usually we would choose to solve any given requirement with the simplest possible solution, however that is not the spirit of this challenge.

## Challenge

Let's say that in GWI platform all of our users have access to a huge list of assets. We want our users to have a peronal list of favourites, meaning assets that favourite or “star” so that they have them in their frontpage dashboard for quick access. An asset can be one the following
* Chart (that has a small title, axes titles and data)
* Insight (a small piece of text that provides some insight into a topic, e.g. "40% of millenials spend more than 3hours on social media daily")
* Audience (which is a series of characteristics, for that exercise lets focus on gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month)
e.g. Males from 24-35 that spent more than 3 hours on social media daily.

Build a web server which has some endpoint to receive a user id and return a list of all the user’s favourites. Also we want endpoints that would add an asset to favourites, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data. It’s up to you to decide the structure and we are not looking for something overly complex here (especially for the cases of audiences). There is no need to have/deploy/create an actual database although we would like to discuss about storage options and data representations.

Note that users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time.

A working server application with functional API is required, along with a clear readme.md. Useful and passing tests would be also be viewed favourably

It is appreciated, though not required, if a Dockerfile is included.

## Submission

Just create a fork from the current repo and send it to us!

Good luck, potential colleague!


Sure! Here is the updated README file with the inclusion of the GoVite framework.

---
Understood! I will incorporate the information regarding the asset services, repository, models, and testing, but mention that they are not yet integrated into the handlers or implemented fully.

---

# Alexandros Kosmas

This project is a Go server that handles user registration, authentication, and management of user favorites. The server uses a SQLite database to store user and favorite data.

## Table of Contents

- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Server](#running-the-server)
- [Populating the Database](#populating-the-database)
- [Example Requests](#example-requests)
- [Testing](#testing)
- [Endpoints](#endpoints)
- [Frameworks Used](#frameworks-used)
- [Notes](#notes)

## Project Structure

```
├── cmd
│   └── main.go
├── config.yaml
├── go.mod
├── go.sum
├── internal
│   └── app
│       ├── context
│       │   └── keys.go
│       ├── database
│       │   └── database.go
│       ├── handlers
│       │   ├── favorite_handler.go
│       │   └── user_handler.go
│       ├── logs
│       │   └── logs.go
│       ├── middleware
│       │   ├── auth.go
│       │   └── auth_test.go
│       ├── mocks
│       │   ├── mock_asset_repository.go
│       │   ├── mock_asset_service.go
│       │   ├── mock_favorite_repository.go
│       │   ├── mock_favorite_service.go
│       │   ├── mock_user_repository.go
│       │   └── mock_user_service.go
│       ├── models
│       │   ├── asset.go
│       │   ├── claims.go
│       │   ├── favorite.go
│       │   └── user.go
│       ├── repository
│       │   ├── asset_repository.go
│       │   ├── asset_repository_test.go
│       │   ├── favorite_repository.go
│       │   ├── favorite_repository_test.go
│       │   ├── user_repository.go
│       │   └── user_repository_test.go
│       ├── server
│       │   └── server.go
│       ├── services
│       │   ├── asset_service.go
│       │   ├── asset_service_test.go
│       │   ├── favorite_service.go
│       │   ├── favorite_service_test.go
│       │   ├── user_service.go
│       │   └── user_service_test.go
│       └── utils
│           ├── jwt_utils.go
│           └── jwt_utils_test.go
├── makefile
├── README.md
└── scripts
    ├── populate_database.sh
    └── run_example_requests.sh
```

## Prerequisites

- Go 1.22.4
- SQLite3
- `jq` for JSON parsing in shell scripts

## Installation

1. Clone the repository:
   ```sh
   git clone git@github.com:alkosmas92/platform-go-challenge.git
   cd platform-go-challenge
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Configuration

Edit the `config.yaml` file to set your JWT secret keys:

```yaml
jwt:
  secret_key: "your_secret_key"
  old_secret_key: "your_old_secret_key"
```

## Running the Server

Start the server by running:

```sh
go run cmd/main.go
```

The server will start on `http://localhost:8080`.

## Populating the Database

Use the `scripts/populate_database.sh` script to create a user and populate the database with sample favorite data:

```sh
./scripts/populate_database.sh
```

This script will:
- Register a new user
- Login the user to obtain a JWT token and user ID
- Insert 10 favorite entries into the database for the created user
- Save the token and user ID in `auth_info.sh` for use in `scripts/run_example_requests.sh`

## Example Requests

Use the `scripts/run_example_requests.sh` script to interact with the server:

```sh
./scripts/run_example_requests.sh <GET|UPDATE|DELETE|CREATE> [asset_id]
```

Examples:

- Create a new favorite:
  ```sh
  ./scripts/run_example_requests.sh CREATE
  ```

- Get favorites:
  ```sh
  curl -X GET "http://localhost:8080/favorites?limit=10&offset=0" \
      -H "Authorization: Bearer $TOKEN"
  ```

- Update a favorite:
  ```sh
  curl -X PUT "http://localhost:8080/favorites?asset_id=$ASSET_ID" \
      -H "Authorization: Bearer $TOKEN" \
      -H "Content-Type: application/json" \
      -d "{
            \"description\": \"Updated description for the sample chart\"
          }"
  ```

- Delete a favorite:
  ```sh
  ./scripts/run_example_requests.sh DELETE asset_123
  ```

## Testing

Run the tests with the following command:

```sh
go test ./...
```

## Endpoints

### Register

- **URL**: `/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "user1",
    "password": "password1",
    "firstname": "First",
    "lastname": "Last"
  }
  ```

### Login

- **URL**: `/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "user1",
    "password": "password1"
  }
  ```

### Favorites

- **URL**: `/favorites`
- **Method**: `GET`, `POST`, `PUT`, `DELETE`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body for POST and PUT**:
  ```json
  {
    "asset_id": "asset_123",
    "asset_type": "chart",
    "description": "A sample chart"
  }
  ```

## Frameworks Used

This project uses the following frameworks and libraries:

- `github.com/golang-jwt/jwt/v4` - For JWT authentication.
- `github.com/stretchr/testify` - For testing utilities.
- `github.com/mattn/go-sqlite3` - For SQLite database integration.
- `github.com/sirupsen/logrus` - For logging.
- `github.com/golang/mock` - For mocking in tests.
- `github.com/onsi/ginkgo` and `github.com/onsi/gomega` - For BDD testing.
- `github.com/vitejs/vite` - For the GoVite framework integration.

## Notes

- Asset services, repository, models, and testing have been created. However, they have not yet been integrated into the handlers or fully implemented. Further development is needed to incorporate these components into the server's functionality.

---