package main

import (
	"accounts/internal/api"
	"accounts/internal/autogen"
	"accounts/internal/config"
	"accounts/internal/db"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg := config.GetConfig()
	
	// Open a database connection
	sqlDB, db, err := db.InitDB()
	if err != nil {
		logrus.Fatalf("init db error: %v", err)
		panic(err)
	}

	defer sqlDB.Close()
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer(db)

	e := echo.New()

	autogen.RegisterHandlers(e, server)

	// And we serve HTTP until the world ends.
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	logrus.Fatal(e.Start(address))
}
