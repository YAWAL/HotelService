package controller

import (
	"net/http"
	"encoding/json"
	"github.com/YAWAL/HotelService/model"
	"path"
	"strings"
)

func RenderJSON(w http.ResponseWriter, response interface{}) {
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// HotelRooms is handler function for "/hotelRooms GET" route
func HotelRooms(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, (model.HotelRoom{}).GetAllHotelRooms())
}

// HotelRoomByNum is handler function for "/hotelRoom/{id:[0-9]+} GET" route
func HotelRoomByNum(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	number := path.Base(url)
	RenderJSON(w, (model.HotelRoom{}).GetHotelRoomByNum(number))
}

// Rents is handler function for "/rents GET" route
func Rents(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, (model.Rent{}).GetAllRents())
}

// Rents is handler function for "/rents GET" route
func ShowRents(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, (model.ShowRent{}).ShowAllRents())
}

// DeleteRent is handler for "/rents/{id} DELETE" route
func DeleteRent(w http.ResponseWriter, r *http.Request) {
	hotelNumber := path.Base(r.URL.Path)
	RenderJSON(w, (model.ShowRent{}).DeleteRent(hotelNumber))

}

// UpdateRents is handler for "/rents/{hotelNum:[0-9]+}/{tenId:[0-9]+} PUT" route
func UpdateRents(w http.ResponseWriter, r *http.Request) {
	tenantId := path.Base(r.URL.Path)
	pth := strings.Split(r.URL.Path, "/")
	hotelNum := pth[len(pth)-2]
	RenderJSON(w, (model.ShowRent{}).UpdateRent(hotelNum, tenantId))
}

// CreateRents is handler for "/rents/{hotelNum:[0-9]+}/{tenId:[0-9]+} POST" route
func CreateRents(w http.ResponseWriter, r *http.Request) {
	tenantId := path.Base(r.URL.Path)
	pth := strings.Split(r.URL.Path, "/")
	hotelNum := pth[len(pth)-2]
	RenderJSON(w, (model.ShowRent{}).CreateRent(hotelNum, tenantId))
}
