package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tanmoy037/mongoapi/router"
)

func main() {
	fmt.Println("This Go-API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listen at port 4000 ...")

}
