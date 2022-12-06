package main

import (
	"BackEnd/database"
	"BackEnd/pkg/connection"
	"BackEnd/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	//env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	//inital DB here
	connection.DatabaseInit()

	//run migration
	database.RunMigration()

	//create route
	r := mux.NewRouter()

	//grouping routes
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//setup static prefix path

	var port = os.Getenv("PORT")
	fmt.Println("server running on port" + port)

	//run server
	http.ListenAndServe(":"+port, (r))
}
