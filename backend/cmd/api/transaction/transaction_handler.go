package transaction

import (
	"encoding/json"
	"net/http"

	"github.com/rohits3m/finance-tracker/internal/models"
	"github.com/rohits3m/finance-tracker/internal/server"
)

type TransactionHandler struct {
	*server.Server
	TransactionModel *models.TransactionModel
}

func NewTransactionHandler(server *server.Server) *TransactionHandler {
	return &TransactionHandler{
		Server:           server,
		TransactionModel: &models.TransactionModel{Db: server.Db},
	}
}

func (handler *TransactionHandler) HandleGetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.TransactionModel.Get()
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, transactions, "")
}

func (handler *TransactionHandler) HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	var data models.CreateTransaction
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := handler.TransactionModel.Create(data)
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, id, "transaction created")
}
