package handler

import (
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/src/service"
)

type Handler struct {
	Core *core.Core
	Ping *PingHandler
	Auth *AuthHandler
	Blog *BlogHandler
}

func NewHandler(
	c *core.Core,
	authService *service.AuthService,
	blogService *service.BlogService,
) *Handler {
	return &Handler{
		Core: c,
		Ping: NewPingHandler(c),
		Auth: NewAuthHandler(c, authService),
		Blog: NewBlogHandler(c, blogService),
	}
}
