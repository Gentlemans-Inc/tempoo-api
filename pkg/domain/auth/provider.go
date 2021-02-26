package auth

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/auth/handler"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	handler.NewAuthHandler,
	)
