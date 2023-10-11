package main

import (
	"fmt"
	"net/http"
	"paytabs-task/controllers"
	"paytabs-task/db"
	"paytabs-task/services"
)

func main() {

	fmt.Println("Loading data ....")
	db := &db.Db{}
	db.LoadDataFromJSONFile()

	services := services.NewService(db)

	controllers := controllers.NewController(services)

	http.HandleFunc("/api/accounts/listing", controllers.GetAllAccounts)
	http.HandleFunc("/api/accounts/transaction", controllers.MakeTransAction)

	fmt.Println("Server is ready running on Port 8080")
	http.ListenAndServe(":8080", nil)

}
