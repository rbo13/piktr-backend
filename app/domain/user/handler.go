package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"piktr/app/response"

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

type UserResponse struct {
	User    *User   `json:"user"`
	Success bool    `json:"success"`
	Err     *string `json:"err"`
	Message string  `json:"message"`
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
	userResp := UserResponse{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.userService.CreateUser(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userResp.User = &user
	userResp.Success = true
	userResp.Message = "User created"
	userResp.Err = nil

	res, err := json.Marshal(userResp)

	if err != nil && res == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = response.JSON(w, res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	userResp := UserResponse{}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := u.userService.FindUserByID(id)

	if err != nil {
		errMessage := err.Error()

		userResp.Success = false
		userResp.Err = &errMessage
		userResp.Message = "User dont exist"

		errResp, err := json.Marshal(userResp)

		if err != nil && errResp == nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = response.JSON(w, errResp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	userResp.User = user
	userResp.Err = nil
	userResp.Success = true
	userResp.Message = "User successfully retrieved"

	res, err := json.Marshal(&userResp)

	if err != nil && res == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = response.JSON(w, res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.FindAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(users)

	if err != nil && res == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = response.JSON(w, res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
