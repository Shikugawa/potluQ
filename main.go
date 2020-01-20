package main

import (
	"context"
	"log"
	"os"
	"time"

	"flag"

	"github.com/Shikugawa/potluq/ent"
	"github.com/Shikugawa/potluq/infra"
	"github.com/Shikugawa/potluq/message"
	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbHost     = os.Getenv("MYSQL_HOST")
	dbPort     = os.Getenv("MYSQL_PORT")
	dbName     = os.Getenv("MYSQL_NAME")
)

func Open() (*ent.Client, error) {
	drv, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
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
	migrate := flag.Bool("auto-migration", false, "migrate database?")
	flag.Parse()

	client, err := Open()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	if *migrate {
		retry := 0
		for {
			if err := client.Schema.Create(context.Background()); err != nil {
				retry += 1
				log.Printf("Failed to migration. Start to retry... %d", retry)
				time.Sleep(time.Second * 5) // wait 5 minutes until mysql is available
				if retry == 5 {
					break
				}
				continue
			}
			log.Printf("migration succeeded. tried %d times.", retry)
			break
		}

	}

	queue := make(chan message.QueueMessage)

	infra.Router(client, &queue)
}
