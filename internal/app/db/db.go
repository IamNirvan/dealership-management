package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/dealership-management/internal/app/structs"
	"github.com/dealership-management/internal/app/util"
	configHandler "github.com/dealership-management/internal/cfg"
	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
	Config     *configHandler.Config
	CloseFunc  func()
}

var (
	database *Database
	dbOnce   sync.Once
)

func Init(config *configHandler.Config) (*Database, error) {
	var dbErr error

	dbOnce.Do(func() {
		connection, openErr := sql.Open("postgres", config.GetConnectionString())
		if openErr != nil {
			database = nil
			dbErr = openErr
			return
		}
		util.LogInfoMsg("successfully initiated database connection")

		// Use the newly obtained connection to ping the database server
		if pingErr := connection.Ping(); pingErr != nil {
			database = nil
			dbErr = pingErr
			return
		}
		util.LogInfoMsg("successfully pinged database")

		// Create the table (if it does not exist)
		if tableCreateErr := createTable(connection, config.Database.Table_name); tableCreateErr != nil {
			database = nil
			dbErr = tableCreateErr
			return
		}
		util.LogInfoMsg("successfully created table (if it does not exist...)")

		// Function to handle database connection closing sequence
		closeFunc := func() {
			connection.Close()
			util.LogInfoMsg("closed database connection")
		}

		database = &Database{
			Connection: connection,
			Config:     config,
			CloseFunc:  closeFunc,
		}
	})

	return database, dbErr
}

func createTable(db *sql.DB, tableName string) error {
	createTableSQL := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
			vehicle_uuid uuid DEFAULT uuid_generate_v1 (),
			manufacturer VARCHAR(255) NOT NULL,
			model VARCHAR(255) NOT NULL,
			colour VARCHAR(255) NOT NULL,
			price VARCHAR(255) NOT NULL,

			PRIMARY KEY(vehicle_uuid)
	)`, tableName)

	_, err := db.Exec(createTableSQL)
	return err
}

func (db *Database) InsertCar(car *structs.Car) error {
	query := fmt.Sprintf(`
	INSERT INTO %s (manufacturer, model, colour, price) VALUES ($1, $2, $3, $4)
	`, db.Config.Database.Table_name)

	util.LogInfoMsg(fmt.Sprintf("Inserting car: %+v", car))
	_, err := db.Connection.Exec(query, car.Manufacturer, car.Model, car.Colour, car.Price)
	return err
}
