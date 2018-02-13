package model

import (
	"time"
)

type Tenant struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type HotelRoom struct {
	Number       int  `json:"number"`
	RoomQuantity int  `json:"room_quantity"`
	IsFree       bool `json:"is_free"`
}

type Rent struct {
	Id           int       `json:"id"`
	HotelRoomNum int       `json:"hotel_room_num"`
	TenantId     int       `json:"tenant_id"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}

// RentResponse represent response (list of ordered rooms with tenant's name) for rendering
type RentResponse struct {
	OrderedRooms map[int]string `json:"ordered_rooms"`
}

func (hr HotelRoom) GetAllHotelRooms() []HotelRoom {
	hotelRooms := make([]HotelRoom, 0)
	//TODO implement GetAllHotelRooms logic
	return hotelRooms
}

func (r Rent) GetAllRents() []Rent {
	rents := make([]Rent, 0)
	//TODO implement GetAllRents logic
	return rents
}
