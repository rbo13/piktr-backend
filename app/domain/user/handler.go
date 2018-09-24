package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Handler defines the list
// of http handlers for handling
// different request
type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService Service
}

// NewHandler returns
// the Handler interface
func NewHandler(userService Service) Handler {
	return &userHandler{
		userService,
	}
}

func (u *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.userService.CreateUser(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)

	if err != nil && response == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		log.Fatalf("Error parsing ID due to: %v", err)
		return
	}

	user, err := u.userService.FindUserByID(id)

	if err != nil {
		log.Fatalf("Error getting user due to: %v", err)
		return
	}

	response, err := json.Marshal(user)

	if err != nil && response == nil {
		log.Fatalf("Error marshaling due to: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)

	if err != nil {
		log.Fatalf("Error writing to write header due to: %v", err)
		return
	}
}

func (u *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	tickets, err := u.userService.FindAllUsers()

	if err != nil {
		log.Fatalf("Error getting list of users due to: %v", err)
		return
	}

	response, err := json.Marshal(tickets)

	if err != nil && response == nil {
		log.Fatalf("Error marshaling due to: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)

	if err != nil {
		log.Fatalf("Error writing to write header due to: %v", err)
		return
	}
}
