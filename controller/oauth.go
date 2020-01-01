package controller

import (
	"encoding/json"
	"net/http"

	"context"

	"fmt"

	"github.com/Shikugawa/potluq/ent"
	"github.com/Shikugawa/potluq/message"
	"github.com/Shikugawa/potluq/service"
)

type OauthController struct {
	oauthService *service.JwtService
	userService  *service.UserService
}

func InitOauthController(client *ent.Client) *OauthController {
	return &OauthController{
		oauthService: &service.JwtService{
			Client: client,
		},
		userService: &service.UserService{
			Client: client,
		},
	}
}

func (oauth *OauthController) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var credential message.Credential
		if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if check := oauth.userService.VerityPassword(context.Background(), &credential); !check {
			http.Error(w, "invalid password", http.StatusUnauthorized)
			return
		}

		user, err := oauth.userService.FindByEmail(context.Background(), credential.Email)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not registered", credential.Email), http.StatusUnauthorized)
			return
		}

		token, err := oauth.oauthService.DefaultJwt(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		w.Write([]byte(token))
		return
	}
}
