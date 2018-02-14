package model

import (
	"time"
	"github.com/YAWAL/HotelService/database"
	"log"
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

// GetAllHotelRooms returns list of all rooms in hotel
func (hr HotelRoom) GetAllHotelRooms() []HotelRoom {
	hotelRooms := make([]HotelRoom, 0)
	query := "SELECT * FROM hotel_rooms ORDER BY number"
	conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
	}
	if rows, err := conn.Raw(query).Rows(); err == nil {
		defer rows.Close()
		for rows.Next() {
			var hotelRoom HotelRoom
			rows.Scan(&hotelRoom.Number, &hotelRoom.RoomQuantity, &hotelRoom.IsFree)
			hotelRooms = append(hotelRooms, hotelRoom)
		}
	}
	return hotelRooms
}

// GetAllRents return list of all ordered rooms
func (r Rent) GetAllRents() []Rent {
	rents := make([]Rent, 0)
	query := "SELECT * FROM rents ORDER BY id"
	conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
	}
	if rows, err := conn.Raw(query).Rows(); err == nil {
		defer rows.Close()
		for rows.Next() {
			var rent Rent
			rows.Scan(&rent.Id, &rent.HotelRoomNum, &rent.TenantId, &rent.StartDate, &rent.EndDate)
			rents = append(rents, rent)
		}
	}
	return rents
}
