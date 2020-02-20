package main

import (
	"log"

	"github.com/renantarouco/clean-todo/api"
	"github.com/renantarouco/clean-todo/core"
	"github.com/renantarouco/clean-todo/database"
	"github.com/spf13/pflag"
)

var (
	dbAddr = pflag.StringP("db", "d", "renan@localhost:4567", "Database address")
)

func main() {
	// CONFIG
	pflag.Parse()
	// CONFIG

	// DI
	db, err := database.NewTasksDB(*dbAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	service := core.NewTasksService(db)
	api := api.NewTasksAPI(service)
	// DI

	if err := api.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
