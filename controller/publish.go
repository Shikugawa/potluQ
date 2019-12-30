package controller

import (
	"encoding/json"
	"net/http"

	"fmt"

	"net/url"

	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/message"
	"github.com/Shikugawa/potraq/service"
)

type PublishController struct {
	Queue       *chan message.QueueMessage
	userService *service.UserService
}

func InitPublishController(client *ent.Client, queue *chan message.QueueMessage) *PublishController {
	return &PublishController{
		Queue: queue,
		userService: &service.UserService{
			Client: client,
		},
	}
}

func (controller *PublishController) EnqueueMusic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var publish message.Publish
		if err := json.NewDecoder(r.Body).Decode(&publish); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if media, err := controller.mediaType(publish.Url); err == nil {
			clubName := r.Header.Get("club_name")
			userName := r.Header.Get("user_name")

			switch media {
			case message.YouTube:
				*controller.Queue <- message.QueueMessage{
					ClubName:  clubName,
					UserName:  userName,
					Url:       publish.Url,
					MediaType: media,
				}
			default:
				http.Error(w, fmt.Sprintf("%s has unsupported media type", publish.Url), http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusAccepted)
		return
	}
}

func (controller *PublishController) mediaType(rawurl string) (message.Media, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return -1, err
	}
	switch u.Hostname() {
	case "youtube.com":
	case "www.youtube.com":
		return message.YouTube, nil
	}
	return -1, nil
}
