package account

import (
	"fmt"

	"github.com/rohits3m/finance-tracker/internal/server"
)

func RegisterAccountRoutes(server *server.Server, base string) {
	handler := NewAccountHandler(server)

	server.Mux.HandleFunc(fmt.Sprintf("GET %s", base), handler.HandleGetAccounts)
	server.Mux.HandleFunc(fmt.Sprintf("POST %s/create", base), handler.HandleCreateAccount)
}
