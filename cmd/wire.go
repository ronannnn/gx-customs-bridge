//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ronannnn/gx-customs-bridge/internal/apis"

	"github.com/google/wire"
)

func InitHttpServer() (*apis.HttpServer, error) {
	wire.Build(wireSet)
	return &apis.HttpServer{}, nil
}
