package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vandit1604/mongoapi/router"
)

func main() {
	fmt.Println("API IS STARTING")
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
}
