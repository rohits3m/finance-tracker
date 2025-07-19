package user

import (
	"encoding/json"
	"net/http"

	"github.com/rohits3m/finance-tracker/internal/models"
	"github.com/rohits3m/finance-tracker/internal/server"
)

type UserHandler struct {
	*server.Server
	UserModel *models.UserModel
}

func NewUserHandler(server *server.Server) *UserHandler {
	return &UserHandler{
		Server:    server,
		UserModel: &models.UserModel{Db: server.Db},
	}
}

func (handler *UserHandler) HandleGetUserByFirebaseId(w http.ResponseWriter, r *http.Request) {
	firebaseId := r.PathValue("firebaseId")
	user, err := handler.UserModel.GetByFirebaseId(firebaseId)
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, user, "")
}

func (handler *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var data models.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := handler.UserModel.Create(data)
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, id, "user created")
}
