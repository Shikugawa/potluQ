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
	service *service.JwtService
}

func InitOauthController(client *ent.Client) *OauthController {
	return &OauthController{
		service: &service.JwtService{
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
		var token string
		var err error

		switch credential.Status {
		case message.JUKEBOX:
			token, err = oauth.service.CreateJukeboxJwt(context.Background(), &credential)
		case message.REQUESTER:
			token, err = oauth.service.CreateRequesterJwt(context.Background(), &credential)
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
