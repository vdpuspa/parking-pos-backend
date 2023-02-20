package api

type CheckInPayload struct {
	VehicleNumber string `json:"VehicleNumber"`
	VehicleType   string `json:"VehicleType"`
}

type CheckOutPayload struct {
	ParkingTicketID int64  `json:"ParkingTicketID"`
	VehicleNumber   string `json:"VehicleNumber"`
	PaymentMethod   string `json:"PaymentMethod"`
	CheckOutTime    string `jsom:"CheckOutTime"`
}
