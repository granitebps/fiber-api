package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/fiber-api/pkg/constants"
	"github.com/granitebps/fiber-api/pkg/utils"
	"github.com/granitebps/fiber-api/src/handler"
	"github.com/spf13/viper"
)

func SetupRoute(a *fiber.App, h *handler.Handler) {
	a.Get("/", func(ctx *fiber.Ctx) error {
		return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, fmt.Sprintf("%s API", viper.GetString(constants.APP_NAME)), nil)
	})

	route := a.Group("/api")

	// V1 Route
	v1Route(route, h)
}

func v1Route(route fiber.Router, h *handler.Handler) {
	v1 := route.Group("/v1")

	v1.Get("/ping", h.Ping.Ping)
}
