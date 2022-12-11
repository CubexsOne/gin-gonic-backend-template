package database

import (
	"fmt"

	"github.com/cubexsone/gin-gonic-backend-template/src/environment"
)

var (
	host     = environment.PG_HOST
	port     = environment.PG_PORT
	user     = environment.PG_USER
	password = environment.PG_PASSWORD
	database = environment.PG_DATABASE
	schema   = environment.PG_SCHEMA
	sslMode  = environment.PG_SSL_MODE
)

func Connect() {
	database := New(host, port, user, password, database, schema, sslMode)
	err := database.Connect()
	if err != nil {
		fmt.Println("Connect to database failed:", err)
	}
}
