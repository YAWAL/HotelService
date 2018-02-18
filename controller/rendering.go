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

// Rents is a handler function for "/rents GET" route
// it shows list with ordered hotel rooms and tenants
func ShowRents(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, (model.ShowRent{}).ShowAllRents())
}

// DeleteRent is a handler for "/rents/{id} DELETE" route
// it deletes item from list with ordered hotel rooms and tenants by hotel number id
func DeleteRent(w http.ResponseWriter, r *http.Request) {
	hotelNumber := path.Base(r.URL.Path)
	RenderJSON(w, (model.ShowRent{}).DeleteRent(hotelNumber))
}

// UpdateRents is a handler for "/rents/{hotelNum:[0-9]+}/{tenId:[0-9]+} PUT" route
// it updates item in list with ordered hotel rooms and tenants
func UpdateRents(w http.ResponseWriter, r *http.Request) {
	tenantId := path.Base(r.URL.Path)
	pth := strings.Split(r.URL.Path, "/")
	hotelNum := pth[len(pth)-2]
	RenderJSON(w, (model.ShowRent{}).UpdateRent(hotelNum, tenantId))
}

// CreateRents is a handler for "/rents/{hotelNum:[0-9]+}/{tenId:[0-9]+} POST" route
// it creates item in list with ordered hotel rooms and tenants
func CreateRents(w http.ResponseWriter, r *http.Request) {
	tenantId := path.Base(r.URL.Path)
	pth := strings.Split(r.URL.Path, "/")
	hotelNum := pth[len(pth)-2]
	RenderJSON(w, (model.ShowRent{}).CreateRent(hotelNum, tenantId))
}
