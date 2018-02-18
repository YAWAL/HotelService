package model

import (
	"github.com/YAWAL/HotelService/database"
	"log"
	"strconv"
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
	Id           int `json:"id"`
	HotelRoomNum int `json:"hotel_room_num"`
	TenantId     int `json:"tenant_id"`
}

type ShowRent struct {
	HotelRoomNum int     `json:"hotel_room_num"`
	Tenant       *Tenant `json:"tenant"`
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

// GetHotelRoomByNum returns info about hotel room with mentioned room number
func (hr HotelRoom) GetHotelRoomByNum(number string) HotelRoom {
	var hotelRoom HotelRoom
	conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
	}
	conn.Where("number = ?", number).Find(&hotelRoom)
	return hotelRoom
}

// GetAllRents return list of all rents
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
			rows.Scan(&rent.Id, &rent.HotelRoomNum, &rent.TenantId)
			rents = append(rents, rent)
		}
	}
	return rents
}

func GetAllRents() []Rent {
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
			rows.Scan(&rent.Id, &rent.HotelRoomNum, &rent.TenantId)
			rents = append(rents, rent)
		}
	}
	return rents
}

func (sr ShowRent) ShowAllRents() []ShowRent {
	rents := GetAllRents()
	log.Println(rents)
	showRents := make([]ShowRent, 0)
	conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
	}
	for _, rent := range rents {
		var showRent ShowRent
		query := "SELECT * FROM tenants WHERE id = ?"
		if row := conn.Raw(query, rent.TenantId).Row(); err == nil {
			var tenant Tenant
			row.Scan(&tenant.Id, &tenant.Name, &tenant.LastName)
			showRent.HotelRoomNum = rent.HotelRoomNum
			showRent.Tenant = &tenant
			showRents = append(showRents, showRent)
		} else {
			log.Fatalf("Error during selecting tenenta has occurred: %v", err)
		}
	}
	return showRents
}

func ShowAllRents() []ShowRent {
	rents := GetAllRents()
	log.Println(rents)
	showRents := make([]ShowRent, 0)
	conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
	}
	for _, rent := range rents {
		var showRent ShowRent
		query := "SELECT * FROM tenants WHERE id = ?"
		if row := conn.Raw(query, rent.TenantId).Row(); err == nil {
			var tenant Tenant
			row.Scan(&tenant.Id, &tenant.Name, &tenant.LastName)
			showRent.HotelRoomNum = rent.HotelRoomNum
			showRent.Tenant = &tenant
			showRents = append(showRents, showRent)
		} else {
			log.Fatalf("Error during selecting tenenta has occurred: %v", err)
		}
	}
	return showRents
}

func (sr ShowRent) DeleteRent(hotelNumber string) []ShowRent {
	deleteQuery := "DELETE FROM rents WHERE hotel_room_num = ?"
	updateQuery := "UPDATE is_free FROM hotel_rooms SET is_free = TRUE"
	conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
	}
	conn.Exec(deleteQuery, hotelNumber)
	conn.Exec(updateQuery)
	return ShowAllRents()
}

func (sr ShowRent) UpdateRent(hotelNumber, tenantId string) []ShowRent {
	updateQuery := "UPDATE rents SET tenant_id = ? WHERE hotel_room_num = ?"
	conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
	}
	conn.Exec(updateQuery, tenantId, hotelNumber)
	return ShowAllRents()
}

func (sr ShowRent) CreateRent(hotelNumber, tenantId string) []ShowRent {
	conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
	}
	hotelNumberInt, err := strconv.Atoi(hotelNumber)
	if err != nil {
		log.Fatal("Error durind conversion of hotelNumber has occurred: ", err)
	}
	tenantIdInt, err := strconv.Atoi(tenantId)
	if err != nil {
		log.Fatal("Error durind conversion of tenantId has occurred: ", err)
	}
	var rent = Rent{HotelRoomNum: hotelNumberInt, TenantId: tenantIdInt}
	conn.Create(&rent)
	return ShowAllRents()
}
