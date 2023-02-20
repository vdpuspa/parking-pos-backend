package repository

import (
	"gorm.io/gorm"
	"parking-pos-backend/domain"
	"parking-pos-backend/src/model"
)

type ParkingRepository struct {
	conn *gorm.DB
}

func NewParkingRepository(conn *gorm.DB) domain.ParkingRepository {
	return &ParkingRepository{
		conn: conn,
	}
}

func (p *ParkingRepository) SaveParkingTicket(ticket model.ParkingTicket) error {
	err := p.conn.Save(&ticket)

	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (p *ParkingRepository) GetParkingTicket(ticket *model.ParkingTicket) error {
	err := p.conn.Where("vehiclenumber = ? and status = ?", ticket.VehicleNumber, ticket.Status).Find(&ticket)

	if err.Error != nil {
		return err.Error
	}

	return nil
}
