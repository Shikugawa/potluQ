package infra

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shikugawa/potluq/controller"
	"github.com/Shikugawa/potluq/ent"
	"github.com/Shikugawa/potluq/external"
	"github.com/Shikugawa/potluq/message"
	"github.com/Shikugawa/potluq/middleware"
	"github.com/gorilla/mux"
)

var (
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")
)

func InitRedisClient() external.RedisHandler {
	return InitRedisHandler(redisHost, redisPort)
}

func Router(dbClient *ent.Client, queue *chan message.QueueMessage) {
	redisClient := InitRedisClient()
	r := mux.NewRouter()

	userController := controller.InitUserController(dbClient)
	oauthController := controller.InitOauthController(dbClient)
	queueController := controller.InitQueueController(dbClient, redisClient, queue)

	authenticator := middleware.InitAuthenticator()
	factory := InitMiddlewareFactory(authenticator.Authenticate)

	r.HandleFunc("/api/user/register", userController.Register)
	r.HandleFunc("/api/auth", oauthController.Auth)
	r.HandleFunc("/api/queue/enqueue", factory.Get(queueController.EnqueueMusic))
	r.HandleFunc("/api/queue/dequeue", factory.Get(queueController.DequeueMusic))

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
	}

	// start 3 workers
	for i := 0; i < 3; i++ {
		worker := InitWorker(redisClient, queue, 5)
		go worker.Start()
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, os.Interrupt)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
