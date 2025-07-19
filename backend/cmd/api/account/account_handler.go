package account

import (
	"encoding/json"
	"net/http"

	"github.com/rohits3m/finance-tracker/internal/models"
	"github.com/rohits3m/finance-tracker/internal/server"
)

type AccountHandler struct {
	*server.Server
	AccountModel *models.AccountModel
}

func NewAccountHandler(server *server.Server) *AccountHandler {
	return &AccountHandler{
		Server:       server,
		AccountModel: &models.AccountModel{Db: server.Db},
	}
}

func (handler *AccountHandler) HandleGetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := handler.AccountModel.Get()
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, accounts, "")
}

func (handler *AccountHandler) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var data models.CreateAccount
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := handler.AccountModel.Create(data)
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, id, "account created")
}
