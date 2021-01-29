package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chaidiryahya/go-simple-api/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	fmt.Println("App Start...")
	port := ":8080"

	router := initRouter()

	log.Println("starting serve on ", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func initRouter() http.Handler {
	router := httprouter.New()

	router.GET("/get", handlers.SampleGet)

	return router
}
