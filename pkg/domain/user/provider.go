package user

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/handler"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/repository"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/services"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	repository.NewUserRepository,
	repository.Set,
	services.NewUserService,
	handler.NewUserHandler,
)
