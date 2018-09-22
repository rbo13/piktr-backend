package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"piktr/app/domain/user"
	"piktr/app/response"
	"piktr/server"

	"piktr/app/database/mysql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	serverAddress = fmt.Sprintf(":%s", os.Getenv("PORT"))
	dbName        = "piktr"
	dns           = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "", "localhost", "3306", "mysql")
)

func main() {
	var userRepo user.Repository

	gormDB, err := setupDatabase(dns)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = gormDB.Debug().Exec("USE " + dbName + ";").Error

	if err != nil {
		log.Fatal(err)
		return
	}

	gormDB.AutoMigrate(&user.User{})

	userRepo = mysql.NewMySQLUserRepository(gormDB)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	srv := server.New(serverAddress, r)

	userService := user.NewUserService(userRepo)
	userHandler := user.NewHandler(userService)

	srv.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8;")
		w.WriteHeader(http.StatusOK)
		response.JSON(w, "Test")
	})

	srv.Router.Get("/user", userHandler.Get)
	srv.Router.Post("/user/create", userHandler.Create)

	srv.Start()

}

func setupDatabase(dns string) (*gorm.DB, error) {
	gormDB, err := gorm.Open("mysql", dns)

	if err != nil {
		return nil, err
	}

	err = gormDB.Debug().Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;").Error

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
