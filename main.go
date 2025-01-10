package main

import (
	"fmt"
	"net/http"
	"noteservice/calc"
	"noteservice/db"
	"noteservice/routes"
)

func main() {

	c := calc.CountVowels("Shubham")
	fmt.Println(c)

	c = calc.CountVowelsOPT("Shubham")
	fmt.Println(c)

	fmt.Println(calc.GenerateStr(10))

	fmt.Println(calc.GenerateStr(10))

	return
	db.InitMongo()

	r := routes.SetUpRoutes()
	fmt.Println("Serving at 8081 ...")
	http.ListenAndServe(":8081", r)
}
