package main

import (
	"fmt"
	"net/http"
	"os"

	"piktr/server"

	"piktr/response"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	serverAddress = fmt.Sprintf(":%s", os.Getenv("PORT"))
	certKey       = "./certs/cert.local.pem"
	privKey       = "./certs/key.local.pem"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	srv := server.New(serverAddress, r)

	srv.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8;")
		w.WriteHeader(http.StatusOK)
		response.JSON(w, "Test")
	})

	srv.Start()
}
