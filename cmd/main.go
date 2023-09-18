package main

import (
	"fmt"

	"github.com/dealership-management/internal/app"
	"github.com/dealership-management/internal/app/db"
	"github.com/dealership-management/internal/app/util"
	configHandler "github.com/dealership-management/internal/cfg"
)

func main() {
	config := configHandler.LoadConfiguration()

	// Initialize the database
	database, err := db.Init(config)
	if err != nil {
		util.LogFatalMsg(fmt.Sprintf("error when initializing the database: %s", err.Error()))
	}
	defer database.CloseFunc()

	application := app.Init(database)
	application.Start()
}
