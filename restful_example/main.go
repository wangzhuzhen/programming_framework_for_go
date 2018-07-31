package main

import (
	"log"
	"net/http"
	"github.com/wangzhuzhen/programming_framework_for_go/restful_example/router"
)

func main() {

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}