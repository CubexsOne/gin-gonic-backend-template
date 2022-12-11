package main

import (
	"github.com/cubexsone/gin-gonic-backend-template/src/database"
	"github.com/cubexsone/gin-gonic-backend-template/src/server"
	"github.com/cubexsone/gin-gonic-backend-template/src/utils/log"
)

func main() {
	log.Info.Println("Server is starting...")
	database.Connect()
	server.Server()
}
