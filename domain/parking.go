package domain

import (
	"parking-pos-backend/src/model"
	"parking-pos-backend/src/model/api"
)

type ParkingUsecase interface {
	CheckIn(payload api.CheckInPayload) error
	CheckOut(payload api.CheckOutPayload) error
	GetParkingTicket(vehicleNumber string) (*model.ParkingTicket, error)
}

type ParkingRepository interface {
	SaveParkingTicket(ticket model.ParkingTicket) error
	GetParkingTicket(ticket *model.ParkingTicket) error
}
