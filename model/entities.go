package model

// Tenant represent tenant entity and table in database
type Tenant struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// HotelRoom represent hotel_room entity and table in database
type HotelRoom struct {
	Number       int  `json:"number"`
	RoomQuantity int  `json:"room_quantity"`
	IsFree       bool `json:"is_free"`
}

// Rent represent rent entity and table in database
type Rent struct {
	Id           int `json:"id"`
	HotelRoomNum int `json:"hotel_room_num"`
	TenantId     int `json:"tenant_id"`
}

// RentToRender is a structure which contain information for rendering in JSON format on web-page
type RentToRender struct {
	HotelRoomNum int     `json:"hotel_room_num"`
	Tenant       *Tenant `json:"tenant"`
}
