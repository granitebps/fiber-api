// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/src/handler"
	"github.com/granitebps/fiber-api/src/repository"
	"github.com/granitebps/fiber-api/src/service"
)

// Injectors from wire.go:

func SetupDependencies(c *core.Core) *handler.Handler {
	blogRepository := repository.NewBlogRepository(c)
	blogService := service.NewBlogService(c, blogRepository)
	handlerHandler := handler.NewHandler(c, blogService)
	return handlerHandler
}
