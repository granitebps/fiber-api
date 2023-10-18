package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type JSONResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func ReturnSuccessResponse(c *fiber.Ctx, code int, msg string, data interface{}) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(code).JSON(JSONResponse{
		Success:   true,
		Message:   msg,
		Data:      data,
		Timestamp: time.Now().UnixMilli(),
	})
}

func ReturnErrorResponse(c *fiber.Ctx, code int, msg string, data interface{}) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(code).JSON(JSONResponse{
		Success:   false,
		Message:   msg,
		Data:      data,
		Timestamp: time.Now().UnixMilli(),
	})
}
