# Requierments

- Go 1.19
- Gitflow

# Installation

1. go get
2. setup .env
3. go run main.go : to migration and start server


# Gitflow
### Installation

- Windows :
1. Download git installer from https://git-scm.com/download/win
2. Install downloaed installer
3. gitflow is in build with git bash
4. gitflow would autimatically would

- Mac OS:
1. brew install git-flow

- Linux:
1. apt-get install git-flow

# Tools And Documentation

- [Gorm](https://gorm.io/).
- [Golang](https://go.dev/).
- [Gin](https://gin-gonic.com/).
- [Environment](https://github.com/joho/godotenv).

# Getting Started

- go run main.go
- go build :
to create compile file: ./dapur-fresh-id 

# Documentation API

### Register
### POST
[https://dapurfresh.herokuapp.com/api/auth/login]

```sh Request Body:
{
    "username":"aldi177",
    "name": "aldi177",
    "phone": "089663652555",
    "password": "aldialdi17"
}

Response Success (status: 201) :
{
    "status": true,
    "message": "Created!",
    "errors": null,
    "data": {
        "id": 5705fdd7-a33d-4647-bef1-9698f423f68c,
        "username": "aldi177",
        "name": "aldi177",
        "phone": "089663652555",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsImV4cCI6MTY1MTgyMDAwMCwiaWF0IjoxNjIwMjg0MDAwLCJpc3MiOiJhZG1pbiJ9.HtnuWlBaevEO3fHAI4McH5W8axvw_3Og47RUI3m9IyI"
    }
}

Error Response (status: 400, 409):
{
    "status": false,
    "message": "Failed to process request",
    "errors": [
        "Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
    ],
    "data": {}
}
{
    "status": false,
    "message": "Failed to process request, duplicate username",
    "errors": [
        ""
    ],
    "data": {}
}
```
### Login
### POST
[https://dapurfresh.herokuapp.com/api/auth/login]

```sh Request Body:
{
    "username":"aldi177",
    "password": "aldialdi17"
}

Response Success (status: 200) :
{
    "status": true,
    "message": "Created!",
    "errors": null,
    "data": {
        "id": 5705fdd7-a33d-4647-bef1-9698f423f68c,
        "username": "aldi177",
        "name": "aldi177",
        "phone": "089663652555",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsImV4cCI6MTY1MTgyMDAwMCwiaWF0IjoxNjIwMjg0MDAwLCJpc3MiOiJhZG1pbiJ9.HtnuWlBaevEO3fHAI4McH5W8axvw_3Og47RUI3m9IyI"
    }
}
```
### Categories
### GET
[https://dapurfresh.herokuapp.com/api/category]

```sh Response Success (status: 200) :

{
    "status": true,
    "message": "Readed!",
    "errors": null,
    "data": {
        "id": 94f55cdc-d3a4-4493-8232-25af8888d216,
        "name": "Sayuran",
    }
}

Error Response (status: 400):

{
    "status": false,
    "message": "Failed to readed",
    "errors": [
        ""
    ],
    "data": {}
}
```
### Categori ID
### GET

[https://dapurfresh.herokuapp.com/api/category/{id}]

```sh Response Success (status: 200) :
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 94f55cdc-d3a4-4493-8232-25af8888d216,
        "name": "Sayuran",
    }
}

Error Response (status: 400):

{
    "status": false,
    "message": "Failed to readed",
    "errors": [
        ""
    ],
    "data": {}
}
```