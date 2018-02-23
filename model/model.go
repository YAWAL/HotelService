package model

import (
	"github.com/YAWAL/HotelService/database"
	"log"
	"strconv"
	"github.com/jinzhu/gorm"
)

var conn, _ = database.DBconnection()

// GetAllRents is a util function which retrieve info about rents from database
func GetAllRents(conn *gorm.DB) ([]Rent, error) {
	rents := make([]Rent, 0)
	query := "SELECT * FROM rents ORDER BY id"
	//conn, err := database.DBconnection()
	//if err != nil {
	//	//	log.Fatal("Error durind connection to db has occurred: ", err)
	//	//	return nil
	//	//}
	if rows, err := conn.Raw(query).Rows(); err == nil {
		defer rows.Close()
		for rows.Next() {
			var rent Rent
			rows.Scan(&rent.Id, &rent.HotelRoomNum, &rent.TenantId)
			rents = append(rents, rent)
		}
	} else {
		log.Fatal("Error durind querying on rents table has occurred: ", err)
		return nil, err
	}
	return rents, nil
}

// ShowAllRents returns list of ordered rooms and their tenants
func (sr RentToRender) ShowAllRents() []RentToRender {
	rents, err := GetAllRents(conn)
	//log.Println(rents)
	showRents := make([]RentToRender, 0)
	//conn, err := database.DBconnection()
	if err != nil {
		log.Fatal("Error durind connection to db has occurred: ", err)
		return nil
	}
	for _, rent := range rents {
		var showRent RentToRender
		query := "SELECT * FROM tenants WHERE id = ?"
		if row := conn.Raw(query, rent.TenantId).Row(); err == nil {
			var tenant Tenant
			row.Scan(&tenant.Id, &tenant.Name, &tenant.LastName)
			showRent.HotelRoomNum = rent.HotelRoomNum
			showRent.Tenant = &tenant
			showRents = append(showRents, showRent)
		} else {
			log.Fatalf("Error during selecting tenenta has occurred: %v", err)
			return nil
		}
	}
	return showRents
}

// ShowAllRents is a util function returns list of ordered rooms and their tenants
func ShowAllRents() []RentToRender {
	rents, err := GetAllRents(conn)
	//log.Println(rents)
	showRents := make([]RentToRender, 0)
	//conn, err := database.DBconnection()
	//if err != nil {
	//	log.Fatal("Error durind connection to db has occurred: ", err)
	//	return nil
	//}
	for _, rent := range rents {
		var showRent RentToRender
		query := "SELECT * FROM tenants WHERE id = ?"
		if row := conn.Raw(query, rent.TenantId).Row(); err == nil {
			var tenant Tenant
			row.Scan(&tenant.Id, &tenant.Name, &tenant.LastName)
			showRent.HotelRoomNum = rent.HotelRoomNum
			showRent.Tenant = &tenant
			showRents = append(showRents, showRent)
		} else {
			log.Fatalf("Error during selecting tenenta has occurred: %v", err)
			return nil
		}
	}
	return showRents
}

// DeleteRent deletes rent cortege and updates hotel_rooms cortage in database
// and returns updated list of ordered rooms and their tenants
func (sr RentToRender) DeleteRent(hotelNumber string) []RentToRender {
	deleteQuery := "DELETE FROM rents WHERE hotel_room_num = ?"
	updateQuery := "UPDATE is_free FROM hotel_rooms SET is_free = true WHERE hotel_room_num = ?"
	//conn, err := database.DBconnection()
	//if err != nil {
	//	log.Fatal("Error durind connection to db has occurred: ", err)
	//}
	conn.Exec(deleteQuery, hotelNumber)
	conn.Exec(updateQuery, hotelNumber)
	return ShowAllRents()
}

// UpdateRent updates info about rent in database and returns updated list of ordered rooms and their tenants
func (sr RentToRender) UpdateRent(hotelNumber, tenantId string) []RentToRender {
	updateQuery := "UPDATE rents SET tenant_id = ? WHERE hotel_room_num = ?"
	//conn, err := database.DBconnection()
	//if err != nil {
	//	log.Fatal("Error durind connection to db has occurred: ", err)
	//}
	conn.Exec(updateQuery, tenantId, hotelNumber)
	return ShowAllRents()
}

// CreateRent create cortege with rent and returns updated list of ordered rooms and their tenants
func (sr RentToRender) CreateRent(hotelNumber, tenantId string) []RentToRender {
	//conn, err := database.DBconnection()
	//if err != nil {
	//	log.Fatal("Error durind connection to db has occurred: ", err)
	//}
	hotelNumberInt, err := strconv.Atoi(hotelNumber)
	if err != nil {
		log.Fatal("Error during conversion of hotelNumber has occurred: ", err)
	}
	tenantIdInt, err := strconv.Atoi(tenantId)
	if err != nil {
		log.Fatal("Error during conversion of tenantId has occurred: ", err)
	}
	var rent = Rent{HotelRoomNum: hotelNumberInt, TenantId: tenantIdInt}
	conn.Create(&rent)
	return ShowAllRents()
}
