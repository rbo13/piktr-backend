package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"piktr/server"

	"piktr/app/repository"
	"piktr/app/response"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	serverAddress = fmt.Sprintf(":%s", os.Getenv("PORT"))
	certKey       = "./certs/cert.local.pem"
	privKey       = "./certs/key.local.pem"
)

var gormDB *gorm.DB

func init() {
	gormDB, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer gormDB.Close()
}

func main() {
	sqlDB := repository.NewSQLDb(gormDB)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	srv := server.New(serverAddress, r)

	log.Print(sqlDB)

	srv.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8;")
		w.WriteHeader(http.StatusOK)
		response.JSON(w, "Test")
	})

	srv.Start()
}
