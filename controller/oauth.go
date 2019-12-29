package controller

import (
	"encoding/json"
	"net/http"

	"context"

	"fmt"

	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/message"
	"github.com/Shikugawa/potraq/service"
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

		var token string

		switch credential.Status {
		case message.JUKEBOX:
			token, err = oauth.oauthService.CreateJukeboxJwt(context.Background(), user)
		case message.REQUESTER:
			token, err = oauth.oauthService.CreateRequesterJwt(context.Background(), user)
		default:
			http.Error(w, fmt.Sprintf("status %s is not supported", credential.Status), http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.Write([]byte(token))
		return
	}
}
