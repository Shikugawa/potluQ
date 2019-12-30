package controller

import (
	"encoding/json"
	"net/http"

	"context"

	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/service"
)

type UserController struct {
	service *service.UserService
}

func InitUserController(c *ent.Client) *UserController {
	return &UserController{
		service: &service.UserService{
			Client: c,
		},
	}
}

func (controller *UserController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user ent.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err := controller.service.FindByEmail(context.Background(), user.Email)
		if err != nil {
			controller.service.CreateUser(context.Background(), &user)
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, "email is already registered", http.StatusBadRequest)
		return
	}
}
