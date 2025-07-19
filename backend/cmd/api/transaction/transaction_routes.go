package transaction

import (
	"fmt"

	"github.com/rohits3m/finance-tracker/internal/server"
)

func RegisterTransactionRoutes(server *server.Server, base string) {
	handler := NewTransactionHandler(server)

	server.Mux.HandleFunc(fmt.Sprintf("GET %s", base), handler.HandleGetTransactions)
	server.Mux.HandleFunc(fmt.Sprintf("POST %s/create", base), handler.HandleCreateTransaction)
}
