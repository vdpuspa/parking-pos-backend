package usecase

import (
	"fmt"
	"parking-pos-backend/domain"
	"parking-pos-backend/src/model"
	"parking-pos-backend/src/model/api"
	"time"
)

type ParkingUsecase struct {
	repo domain.ParkingRepository
}

func NewParkingUsecase(repo domain.ParkingRepository) domain.ParkingUsecase {
	return &ParkingUsecase{
		repo: repo,
	}
}

func (p *ParkingUsecase) CheckIn(payload api.CheckInPayload) error {
	utcLocation := time.UTC
	now := time.Now().In(utcLocation)

	ticket := model.ParkingTicket{
		VehicleNumber: payload.VehicleNumber,
		Status:        model.Parking.String(),
		CheckInTime:   now,
	}

	err := p.repo.SaveParkingTicket(ticket)
	if err != nil {
		return err
	}

	return nil
}

func (p *ParkingUsecase) CheckOut(payload api.CheckOutPayload) error {
	utcLocation := time.UTC
	now := time.Now().In(utcLocation)

	ticket := model.ParkingTicket{
		ParkingTicketID: payload.ParkingTicketID,
		VehicleNumber:   payload.VehicleNumber,
		PaymentMethod:   payload.PaymentMethod,
		CheckOutTime:    now,
		Status:          model.Completed.String(),
	}

	err := p.repo.SaveParkingTicket(ticket)
	if err != nil {
		return err
	}

	return nil
}

func (p *ParkingUsecase) GetParkingTicket(vehicleNumber string) (*model.ParkingTicket, error) {
	fmt.Println(vehicleNumber)
	ticket := &model.ParkingTicket{
		VehicleNumber: vehicleNumber,
		Status:        model.Parking.String(),
	}

	err := p.repo.GetParkingTicket(ticket)
	if err != nil {
		return nil, err
	}
	fmt.Println(ticket)
	return ticket, nil
}
