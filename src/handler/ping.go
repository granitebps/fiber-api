package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/pkg/utils"
)

type PingHandler struct {
	Core *core.Core
}

func NewPingHandler(c *core.Core) *PingHandler {
	return &PingHandler{
		Core: c,
	}
}

// Ping
// @summary Ping to the service
// @description Do a ping to service just make sure service is working normally.
// @tags tools
// @produce json
// @router /api/v1/ping [get]
// @success 200 {object} utils.JSONResponse
func (h *PingHandler) Ping(c *fiber.Ctx) error {
	return utils.ReturnSuccessResponse(c, fiber.StatusOK, "OK", nil)
}
