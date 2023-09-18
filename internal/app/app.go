package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dealership-management/internal/app/db"
	"github.com/dealership-management/internal/app/structs"
	"github.com/dealership-management/internal/app/util"
)

type App struct {
	Database *db.Database
}

func Init(db *db.Database) *App {
	return &App{
		Database: db,
	}
}

func (app *App) Start() {
	var command, manufacturer, model, colour string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter X[exit] or I[insert]: ")
		fmt.Scanln(&command)

		if strings.ToLower(command) == "x" {
			break
		}

		fmt.Print("\nEnter manufacturer: ")
		manufacturer, _ = reader.ReadString('\n')
		fmt.Print("Enter model: ")
		model, _ = reader.ReadString('\n')
		fmt.Print("Enter colour: ")
		colour, _ = reader.ReadString('\n')
		fmt.Print("Enter price: ")
		price, _ := reader.ReadString('\n')

		car := &structs.Car{
			Manufacturer: strings.TrimSpace(manufacturer),
			Model:        strings.TrimSpace(model),
			Colour:       strings.TrimSpace(colour),
			Price:        strings.TrimSpace(price),
		}

		if err := app.Database.InsertCar(car); err != nil {
			util.LogErrorMsg(fmt.Sprintf("failed to insert car with error: %s", err.Error()))
		}
	}
}
