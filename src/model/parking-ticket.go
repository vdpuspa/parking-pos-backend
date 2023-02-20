package model

import "time"

type ParkingTicket struct {
	ParkingTicketID int64     `gorm:"column:parkingticketid;type:bigserial;autoIncrement;primary_key"`
	VehicleNumber   string    `json:"VehicleNumber" gorm:"column:vehiclenumber;type:varchar(10);index"`
	VehicleType     string    `json:"VehicleType" gorm:"column:vehicletype;type:varchar(20)"`
	CheckInTime     time.Time `json:"CheckInTime" gorm:"column:checkintime"`
	CheckOutTime    time.Time `json:"CheckOutTime" gorm:"column:checkouttime"`
	PaymentMethod   string    `json:"PaymentMethod" gorm:"column:paymentmethod;type:varhcar(50)"`
	Status          string    `json:"Status" gorm:"column:status;type:varchar(3)"`
}

func (ParkingTicket) TableName() string {
	return "parkingticket"
}

type StatusEnum int

const (
	Parking StatusEnum = iota
	Completed
)

var StatusEnumString = [...]string{
	"PRK",
	"CMP",
}

func (d StatusEnum) String() string {
	if Parking <= d && d <= Completed {
		return StatusEnumString[d]
	}
	return "Invalid Enum"
}
