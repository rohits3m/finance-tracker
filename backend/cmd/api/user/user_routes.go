package user

import (
	"fmt"

	"github.com/rohits3m/finance-tracker/internal/server"
)

func RegisterUserRoutes(server *server.Server, base string) {
	handler := NewUserHandler(server)

	server.Mux.HandleFunc(fmt.Sprintf("GET %s/{firebaseId}", base), handler.HandleGetUserByFirebaseId)
	server.Mux.HandleFunc(fmt.Sprintf("POST %s", base), handler.HandleCreateUser)
}
