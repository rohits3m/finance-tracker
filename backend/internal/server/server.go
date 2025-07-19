package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServerResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type ServerConfig struct {
	Port  string
	Env   string
	DbStr string
}

type Server struct {
	Mux    *http.ServeMux
	Logger *slog.Logger
	Db     *pgxpool.Pool
	Config *ServerConfig
}

func NewServer(config ServerConfig) (*Server, error) {
	conn, err := pgxpool.New(context.Background(), config.DbStr)
	if err != nil {
		return nil, err
	}

	return &Server{
		Mux:    http.NewServeMux(),
		Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})),
		Db:     conn,
		Config: &config,
	}, nil
}

func (server *Server) Run() error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", server.Config.Port),
		Handler:      server.Mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(server.Logger.Handler(), slog.LevelError),
	}

	err := srv.ListenAndServe()
	return err
}

func (server *Server) SuccessResponse(w http.ResponseWriter, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ServerResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func (server *Server) FailureResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(ServerResponse{
		Success: true,
		Data:    nil,
		Message: message,
	})
}
