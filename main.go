package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/fiber-api/config"
	"github.com/granitebps/fiber-api/pkg/constants"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/scheduler"
	"github.com/granitebps/fiber-api/src/middleware"
	"github.com/granitebps/fiber-api/src/route"
	"github.com/spf13/viper"
)

// @title Fiber API
// @version 1.0
// @description This is a Fiber API Doc
// @contact.name Granite Bagas
// @contact.url https://granitebps.com
// @contact.email granitebagas28@gmail.com
// @license.name MIT
// @BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Load ENV and setup some config
	config.SetupConfig(".env")

	// Initiate Fiber
	app := fiber.New(config.FiberConfig())

	// Setup core package
	c := core.SetupCore()

	// Setup middleware
	middleware.SetupMiddleware(app, c)

	// Setup Dependency Injection
	h := SetupDependencies(c)

	// Setup route
	route.SetupRoute(app, h)

	// Setup scheduler
	scheduler.SetupScheduler(c)

	startServerWithGracefulShutdown(app)
}

func startServerWithGracefulShutdown(app *fiber.App) {
	PORT := viper.GetString(constants.APP_PORT)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(PORT); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	log.Println("Fiber was successful shutdown.")
}

// func startServerWithGracefulShutdown(app *fiber.App) {
// 	PORT := viper.GetString(constants.APP_PORT)

// 	// Create channel for idle connections.
// 	idleConnsClosed := make(chan struct{})

// 	go func() {
// 		sigint := make(chan os.Signal, 1)
// 		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
// 		<-sigint

// 		// Received an interrupt signal, shutdown.
// 		if err := app.Shutdown(); err != nil {
// 			// Error from closing listeners, or context timeout:
// 			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
// 		}

// 		close(idleConnsClosed)
// 	}()

// 	// Run server.
// 	if err := app.Listen(PORT); err != nil {
// 		log.Printf("Oops... Server is not running! Reason: %v", err)
// 	}

// 	<-idleConnsClosed
// }
