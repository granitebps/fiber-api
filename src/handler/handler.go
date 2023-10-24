package handler

import (
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/src/service"
)

type Handler struct {
	Core *core.Core
	Ping *PingHandler
	Blog *BlogHandler
}

func NewHandler(c *core.Core, blogService *service.BlogService) *Handler {
	return &Handler{
		Core: c,
		Ping: NewPingHandler(c),
		Blog: NewBlogHandler(c, blogService),
	}
}
