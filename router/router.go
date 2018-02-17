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
	api.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to hotel service api!")})
	api.HandleFunc("/hotelRooms", controller.HotelRooms).Methods("GET")
	api.HandleFunc("/hotelRoom/{number:[0-9]+}", controller.HotelRoomByNum).Methods("GET")
	api.HandleFunc("/rentsRoom", controller.Rents).Methods("GET")

	api.HandleFunc("/rents", controller.ShowRents).Methods("GET")
	//api.HandleFunc("/rents", controller.CreateRents).Methods("POST")
	//api.HandleFunc("/rents/{id}", controller.UpdateRents).Methods("UPDATE")
	//api.HandleFunc("/rents/{id}", controller.DdeleteRents).Methods("DELETE")

}
