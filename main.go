package main

import (
	"net/http"
	"log"
	"github.com/YAWAL/HotelService/router"
)

func main() {

	r := router.Router

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	log.Print(srv.ListenAndServe())

}
