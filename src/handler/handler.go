package handler

import (
	"github.com/granitebps/fiber-api/pkg/core"
)

type Handler struct {
	Core *core.Core
	Ping *PingHandler
}

func NewHandler(c *core.Core) *Handler {
	return &Handler{
		Core: c,
		Ping: NewPingHandler(c),
	}
}
