package infra

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shikugawa/potraq/controller"
	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/message"
	"github.com/Shikugawa/potraq/middleware"
	"github.com/gorilla/mux"
)

func Router(dbClient *ent.Client, redisClient *RedisHandler, queue *chan message.QueueMessage) {
	r := mux.NewRouter()

	userController := controller.InitUserController(dbClient)
	oauthController := controller.InitOauthController(dbClient)
	publishController := controller.InitPublishController(dbClient, queue)

	authenticator := middleware.InitAuthenticator()
	factory := InitMiddlewareFactory(authenticator.Authenticate)

	r.HandleFunc("/api/user/register", userController.Register)
	r.HandleFunc("/api/auth", oauthController.Auth)
	r.HandleFunc("/api/queue/enqueue", factory.Get(publishController.EnqueueMusic))
	r.HandleFunc("/api/queue/dequeue", nil)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
