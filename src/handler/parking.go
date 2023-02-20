package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"parking-pos-backend/domain"
	"parking-pos-backend/src/model/api"
	"parking-pos-backend/utils"
)

type ParkingHandler struct {
	usecase domain.ParkingUsecase
}

func NewParkingHandler(app *fiber.App, us domain.ParkingUsecase) {
	handler := &ParkingHandler{
		usecase: us,
	}

	app.Post("checkIn", handler.CheckIn)
	app.Post("checkOut", handler.CheckOut)
	app.Get("getParkingTicket/:vehicleNumber", handler.GetParkingTicket)

}

func (p *ParkingHandler) CheckIn(c *fiber.Ctx) error {
	payload := api.CheckInPayload{}

	err := c.BodyParser(&payload)
	if err != nil {
		return utils.GenerateResponse(c, false, err, "Invalid Request", http.StatusBadRequest)
	}

	err = p.usecase.CheckIn(payload)
	if err != nil {
		return utils.GenerateResponse(c, false, nil, "Failed to check in ticket", http.StatusBadRequest)
	}
	return utils.GenerateResponse(c, true, nil, "Parking Ticket Successfully Created", http.StatusOK)
}

func (p *ParkingHandler) CheckOut(c *fiber.Ctx) error {
	payload := api.CheckOutPayload{}

	err := c.BodyParser(&payload)
	if err != nil {
		return utils.GenerateResponse(c, false, err, "Invalid Request", http.StatusBadRequest)
	}

	err = p.usecase.CheckOut(payload)
	if err != nil {
		return utils.GenerateResponse(c, false, nil, "Failed to check out ticket", http.StatusBadRequest)
	}
	return utils.GenerateResponse(c, true, nil, "Parking Ticket Successfully Created", http.StatusOK)
}

func (p *ParkingHandler) GetParkingTicket(c *fiber.Ctx) error {
	vehicleNumber := c.Params("vehicleNumber")
	if vehicleNumber == "" {
		return utils.GenerateResponse(c, false, nil, "Invalid Param", http.StatusBadRequest)
	}
	fmt.Println(vehicleNumber)
	res, err := p.usecase.GetParkingTicket(vehicleNumber)
	if err != nil {
		return utils.GenerateResponse(c, false, nil, "Failed to retrieve ticket", http.StatusBadRequest)
	}
	return utils.GenerateResponse(c, true, res, "Parking Ticket Successfully Retrieved", http.StatusOK)
}
