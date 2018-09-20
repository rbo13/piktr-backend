package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"piktr/server"

	"piktr/app/db"
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
var dbName = "piktr"

var dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "", "localhost", "3306", "mysql")

func init() {
	gormDB, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	newDB := db.New(gormDB)
	
	err = db.Setup(gormDB, dbName)
	if err != nil {
		log.Fatal(err)
	}

	newDB.MigrateTables()

	gormDB.Close()
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	srv := server.New(serverAddress, r)

	sqlDB := repository.NewSQLDb(gormDB)
	log.Print(sqlDB)

	srv.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8;")
		w.WriteHeader(http.StatusOK)
		response.JSON(w, "Test")
	})

	srv.Start()
}
