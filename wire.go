//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/src/handler"
)

func SetupDependencies(c *core.Core) *handler.Handler {
	wire.Build(
		handler.NewHandler,
	)

	return &handler.Handler{}
}
