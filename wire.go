//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/src/handler"
	"github.com/granitebps/fiber-api/src/repository"
	"github.com/granitebps/fiber-api/src/service"
)

func SetupDependencies(c *core.Core) *handler.Handler {
	wire.Build(
		repository.NewBlogRepository,
		repository.NewUserRepository,
		service.NewBlogService,
		service.NewAuthService,

		handler.NewHandler,
	)

	return &handler.Handler{}
}
