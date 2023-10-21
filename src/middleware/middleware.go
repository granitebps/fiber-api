package middleware

import (
	"time"

	_ "github.com/granitebps/fiber-api/docs"

	"github.com/gofiber/contrib/fibernewrelic"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/pkg/utils"
)

func SetupMiddleware(a *fiber.App, c *core.Core) {
	// Fiber Middleware
	a.Use(logger.New())
	a.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Swagger
	// We need to put swagger middleware in here to prevent collision with security middleware
	a.Get("/swagger/*", swagger.HandlerDefault)

	// Another Fiber Middleware
	a.Use(etag.New())
	a.Use(compress.New())
	a.Use(cors.New())
	a.Use(requestid.New())
	a.Use(helmet.New())
	a.Use(limiter.New(limiter.Config{
		Max:               100,
		Expiration:        1 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			return utils.ReturnErrorResponse(c, fiber.StatusTooManyRequests, "Too many requests.", nil)
		},
	}))

	// Sentry
	a.Use(fibersentry.New(fibersentry.Config{
		Repanic: true,
	}))

	// Newrelic
	a.Use(fibernewrelic.New(fibernewrelic.Config{
		Application: c.Newrelic,
	}))

	// Uncomment these code if you want to implement https://docs.gofiber.io/api/middleware/monitor
	// a.Get("/metrics", monitor.New(monitor.Config{
	// 	Title: fmt.Sprintf("%s Monitor", c.AppName),
	// }))
}
