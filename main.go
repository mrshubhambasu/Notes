package main

import (
	"fmt"
	"net/http"
	"noteservice/db"
	"noteservice/routes"
)

func main() {

	db.InitMongo()

	r := routes.SetUpRoutes()
	fmt.Println("Serving at 8081 ...")
	http.ListenAndServe(":8081", r)
}
