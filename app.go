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

	router.NotFound = http.HandlerFunc(handlers.NotFound)

	router.GET("/", handlers.Home)
	router.GET("/get", handlers.SampleGet)
	router.GET("/get_with_route_param/:user_id", handlers.SampleGetWithRouteParam)
	router.GET("/get_with_query_param", handlers.SampleGetWithQueryParam)
	router.POST("/post_with_json_body", handlers.SamplePostwithJsonBody)

	log.Println("Router : ready")
	return router
}
