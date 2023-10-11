package main

import (
	"fmt"
	"net/http"
	"paytabs-task/controllers"
	"paytabs-task/db"
	"paytabs-task/services"
	"time"
)

func main() {
	fmt.Print("Loading data ....")
	db := &db.Db{}
	db.LoadDataFromJSONFile()

	services := services.NewService(db)

	controllers := controllers.NewController(services)

	http.HandleFunc("/api/accounts/listing", controllers.GetAllAccounts)
	http.HandleFunc("/api/accounts/transaction", controllers.MakeTransAction)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error:", err)
	}
	time.Sleep(200 * time.Millisecond)

	fmt.Print("Server is ready ")

}
