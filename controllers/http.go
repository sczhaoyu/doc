package controllers

import (
	"log"
	"net/http"
)

func StartHttp() {
	steupRoutes()
	log.Println("server start port 9088")
	http.ListenAndServe(":9088", nil)
}
