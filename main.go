//go:generate goagen bootstrap -d github.com/nubunto/autobus-web/design

package main

import (
	"database/sql"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/nubunto/autobus-web/app"

	_ "github.com/lib/pq"
)

func main() {
	// Create service
	service := goa.New("autobus-web")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	db, err := sql.Open("postgres", "user=postgres password=postgres sslmode=disable dbname=autobus")
	if err != nil {
		service.LogError("startup", "err", err)
		return
	}

	c := NewGPSController(service, db)
	app.MountGPSController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
