// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"

	"github.com/Shikugawa/potraq/ent/device"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleClub() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the club's edges.
	m0 := client.Music.
		Create().
		SaveX(ctx)
	log.Println("music created:", m0)
	d1 := client.Device.
		Create().
		SetRole(device.RoleJukebox).
		SaveX(ctx)
	log.Println("device created:", d1)

	// create club vertex with its edges.
	c := client.Club.
		Create().
		SetName("string").
		AddMusic(m0).
		AddDevice(d1).
		SaveX(ctx)
	log.Println("club created:", c)

	// query edges.
	m0, err = c.QueryMusic().First(ctx)
	if err != nil {
		log.Fatalf("failed querying music: %v", err)
	}
	log.Println("music found:", m0)

	d1, err = c.QueryDevice().First(ctx)
	if err != nil {
		log.Fatalf("failed querying device: %v", err)
	}
	log.Println("device found:", d1)

	// Output:
}
func ExampleDevice() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the device's edges.
	u1 := client.User.
		Create().
		SetName("string").
		SetEmail("string").
		SetPassword("string").
		SaveX(ctx)
	log.Println("user created:", u1)

	// create device vertex with its edges.
	d := client.Device.
		Create().
		SetRole(device.RoleJukebox).
		AddUser(u1).
		SaveX(ctx)
	log.Println("device created:", d)

	// query edges.

	u1, err = d.QueryUser().First(ctx)
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}
	log.Println("user found:", u1)

	// Output:
}
func ExampleMusic() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the music's edges.

	// create music vertex with its edges.
	m := client.Music.
		Create().
		SaveX(ctx)
	log.Println("music created:", m)

	// query edges.

	// Output:
}
func ExampleUser() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the user's edges.

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetName("string").
		SetEmail("string").
		SetPassword("string").
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.

	// Output:
}
