package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/fiber-api/pkg/constants"
	"github.com/granitebps/fiber-api/pkg/utils"
	"github.com/granitebps/fiber-api/src/handler"
	"github.com/granitebps/fiber-api/src/middleware"
	"github.com/spf13/viper"
)

func SetupRoute(a *fiber.App, h *handler.Handler) {
	a.Get("", func(ctx *fiber.Ctx) error {
		return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, fmt.Sprintf("%s API", viper.GetString(constants.APP_NAME)), nil)
	})

	route := a.Group("api")

	// V1 Route
	v1Route(route, h)
}

func v1Route(route fiber.Router, h *handler.Handler) {
	v1 := route.Group("v1")

	v1.Get("ping", h.Ping.Ping)

	// Auth routes
	auth := v1.Group("auth")
	auth.Post("register", h.Auth.Register)
	auth.Post("login", h.Auth.Login)
	auth.Get("me", middleware.Private(), h.Auth.Me) // You can register JWT middleware to spesific route

	// Blog routes
	blog := v1.Group("blogs")
	blog.Get("", middleware.Private(), h.Blog.Index)
	blog.Get(":id", h.Blog.Show)
	blog.Post("", h.Blog.Store)
	blog.Put(":id", h.Blog.Update)
	blog.Delete(":id", h.Blog.Destroy)
}
