package infra

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/interface/middleware"
	"github.com/gorilla/mux"
)

func Router(dbClient *ent.Client, redisClient *middleware.RedisHandler) {
	r := mux.NewRouter()
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
