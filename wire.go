// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/api/router"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather"
	"github.com/google/wire"
)

func initializeServer() (*router.Server, error) {

	wire.Build(user.Set, weather.Set, router.NewServer)

	return &router.Server{}, nil
}

