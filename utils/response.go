package utils

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func GenerateResponse(c *fiber.Ctx, success bool, data interface{}, message string, code int) error {
	res := &BaseResponse{
		Success: success,
		Result:  data,
		Message: message,
	}
	jsonByte, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return GenerateRawResponseByte(c, success, code, jsonByte)
}

func GenerateRawResponseByte(cFiber *fiber.Ctx, sucess bool, status int, data []byte) error {
	c := cFiber.Context()
	c.Response.SetBodyRaw(data)
	c.Response.Header.SetContentType("application/json")
	return nil
}
