package main

import (
	"github.com/cubexsone/gin-gonic-backend-template/src/database"
	"github.com/cubexsone/gin-gonic-backend-template/src/server"
)

func main() {
	database.Connect()
	server.Server()
}
