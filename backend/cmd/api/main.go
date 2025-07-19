package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rohits3m/finance-tracker/cmd/api/account"
	"github.com/rohits3m/finance-tracker/cmd/api/transaction"
	"github.com/rohits3m/finance-tracker/cmd/api/user"
	"github.com/rohits3m/finance-tracker/internal/server"
)

func main() {
	// Initializing the .env file
	godotenv.Load()
	config := server.ServerConfig{
		Port:  os.Getenv("PORT"),
		Env:   os.Getenv("ENV"),
		DbStr: os.Getenv("DB_URL"),
	}

	// New instance of the server
	server, err := server.NewServer(config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Registering the routes
	user.RegisterUserRoutes(server, "/api/v1/user")
	account.RegisterAccountRoutes(server, "/api/v1/account")
	transaction.RegisterTransactionRoutes(server, "/api/v1/transaction")

	// Starting the server
	server.Logger.Info("Server listening at", "port", server.Config.Port, "env", server.Config.Env)
	if err := server.Run(); err != nil {
		server.Logger.Error(err.Error())
	}
}
