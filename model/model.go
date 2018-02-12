package model

import "time"

type Tenant struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type HotelRoom struct {
	Number       byte `json:"number"`
	RoomQuantity byte `json:"room_quantity"`
	IsFree       bool `json:"is_free"`
}

type Rent struct {
	Id           int       `json:"id"`
	HotelRoomNum byte      `json:"hotel_room_num"`
	Tenants      []Tenant  `json:"tenants"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}
