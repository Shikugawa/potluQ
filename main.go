package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Shikugawa/potluq/ent"
	"github.com/Shikugawa/potluq/infra"
	"github.com/Shikugawa/potluq/message"
	"github.com/facebookincubator/ent/dialect/sql"
)

var (
	dbUser     = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbHost     = os.Getenv("MYSQL_HOST")
	dbPort     = os.Getenv("MYSQL_PORT")
)

func Open() (*ent.Client, error) {
	drv, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort)
	if err != nil {
		return nil, err
	}
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv)), nil
}

func main() {
	client, err := Open()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Auto migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	queue := make(chan message.QueueMessage)

	infra.Router(client, &queue)
}
