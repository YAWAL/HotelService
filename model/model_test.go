package model

import (
	"reflect"
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/stretchr/testify/assert"
)

func mockDB() (sqlmock.Sqlmock, *gorm.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Error durind mocking database: ", err)
		return nil, nil, err
	}
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		log.Fatal("Error durind connection to postgres database: ", err)
		return nil, nil, err
	}
	gormDB.LogMode(true)
	return mock, gormDB, err
}

func TestGetAllRents(t *testing.T) {
	mock, _, _ := mockDB()
	var fieldsNames = []string{"id", "hotel_room_num", "tenant_id"}
	rows := sqlmock.NewRows(fieldsNames)
	rentCortages := Rent{Id:11,HotelRoomNum:22, TenantId:33}
	rows = rows.AddRow(rentCortages)
	expectedRents := []Rent{rentCortages}
	mock.ExpectQuery("SELECT * FROM rents").WillReturnRows(rows)
	returnedRents := GetAllRents()
	assert.Equal(t, returnedRents,expectedRents)

	tests := []struct {
		name string
		want []Rent
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllRents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllRents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRentToRender_ShowAllRents(t *testing.T) {
	type fields struct {
		HotelRoomNum int
		Tenant       *Tenant
	}
	tests := []struct {
		name   string
		fields fields
		want   []RentToRender
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := RentToRender{
				HotelRoomNum: tt.fields.HotelRoomNum,
				Tenant:       tt.fields.Tenant,
			}
			if got := sr.ShowAllRents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RentToRender.ShowAllRents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowAllRents(t *testing.T) {
	tests := []struct {
		name string
		want []RentToRender
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShowAllRents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowAllRents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRentToRender_DeleteRent(t *testing.T) {
	type fields struct {
		HotelRoomNum int
		Tenant       *Tenant
	}
	type args struct {
		hotelNumber string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []RentToRender
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := RentToRender{
				HotelRoomNum: tt.fields.HotelRoomNum,
				Tenant:       tt.fields.Tenant,
			}
			if got := sr.DeleteRent(tt.args.hotelNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RentToRender.DeleteRent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRentToRender_UpdateRent(t *testing.T) {
	type fields struct {
		HotelRoomNum int
		Tenant       *Tenant
	}
	type args struct {
		hotelNumber string
		tenantId    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []RentToRender
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := RentToRender{
				HotelRoomNum: tt.fields.HotelRoomNum,
				Tenant:       tt.fields.Tenant,
			}
			if got := sr.UpdateRent(tt.args.hotelNumber, tt.args.tenantId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RentToRender.UpdateRent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRentToRender_CreateRent(t *testing.T) {
	type fields struct {
		HotelRoomNum int
		Tenant       *Tenant
	}
	type args struct {
		hotelNumber string
		tenantId    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []RentToRender
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := RentToRender{
				HotelRoomNum: tt.fields.HotelRoomNum,
				Tenant:       tt.fields.Tenant,
			}
			if got := sr.CreateRent(tt.args.hotelNumber, tt.args.tenantId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RentToRender.CreateRent() = %v, want %v", got, tt.want)
			}
		})
	}
}
