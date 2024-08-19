//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/icyre/icyre/internal/config"
)

func New() (*App, error) {
	panic(wire.Build(
		newApp,
		wire.NewSet(config.New),
	))
}
