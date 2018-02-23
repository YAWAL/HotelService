package router

import (
	"github.com/gorilla/mux"
	"github.com/YAWAL/HotelService/controller"
	"net/http"
	"fmt"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	api := Router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to hotel service api!")
	})
	api.HandleFunc("/rents", controller.ShowRents).Methods("GET")
	api.HandleFunc("/rents/{hotelNum:[0-9]+}/{tenId:[0-9]+}", controller.CreateRents).Methods("POST")
	api.HandleFunc("/rents/{hotelNum:[0-9]+}/{tenId:[0-9]+}", controller.UpdateRents).Methods("PUT")
	api.HandleFunc("/rents/{id}", controller.DeleteRent).Methods("DELETE")
}
