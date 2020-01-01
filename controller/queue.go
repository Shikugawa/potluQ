package controller

import (
	"encoding/json"
	"net/http"

	"fmt"

	"net/url"

	"github.com/Shikugawa/potluq/ent"
	"github.com/Shikugawa/potluq/external"
	"github.com/Shikugawa/potluq/message"
	"github.com/Shikugawa/potluq/service"
)

type QueueController struct {
	Queue           *chan message.QueueMessage
	consumerService *service.Consumer
	userService     *service.UserService
}

func InitQueueController(client *ent.Client, redisHandler external.RedisHandler, queue *chan message.QueueMessage) *QueueController {
	return &QueueController{
		Queue: queue,
		consumerService: &service.Consumer{
			Handler: redisHandler,
		},
		userService: &service.UserService{
			Client: client,
		},
	}
}

func (controller *QueueController) EnqueueMusic(w http.ResponseWriter, r *http.Request) {
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

func (controller *QueueController) DequeueMusic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		clubName := r.Header.Get("club_name")
		music, err := controller.consumerService.Consume(clubName)
		w.WriteHeader(http.StatusOK)

		if err != nil {
			errMsg := message.Error{
				Message: err.Error(),
			}
			byteMsg, _ := json.Marshal(errMsg)
			w.Write(byteMsg)
			return
		}

		byteMsg, _ := json.Marshal(music)
		w.Write(byteMsg)
		return
	}
}

func (controller *QueueController) mediaType(rawurl string) (message.Media, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return -1, err
	}
	switch u.Hostname() {
	case "www.youtube.com", "youtube.com":
		return message.YouTube, nil
	}
	return -1, nil
}
