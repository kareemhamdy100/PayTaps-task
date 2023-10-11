package test

import (
	"fmt"
	"os"
	"paytabs-task/controllers"
	"paytabs-task/db"
	"paytabs-task/services"
	"testing"
)

type Deps struct {
	db          *db.Db
	services    *services.Services
	controllers *controllers.Controllers
}

var App *Deps = nil

func setup() {
	// Perform setup actions before running tests
	fmt.Println("Setting Up")
	db := &db.Db{}
	db.LoadDataFromJSONFile()

	services := services.NewService(db)

	Contorller := controllers.NewController(services)

	App = &Deps{
		db:          db,
		services:    services,
		controllers: Contorller,
	}
}

func teardown() {
	// Perform teardown actions after running tests
	fmt.Println("tear Down ")
}

func TestMain(m *testing.M) {
	setup() // Setup before running tests

	// Run the tests
	exitCode := m.Run()

	teardown() // Teardown after running tests

	os.Exit(exitCode)
}
