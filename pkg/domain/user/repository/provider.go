package repository

import (
	"github.com/Mangaba-Labs/tempoo-api/database"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	database.NewDatabase,
	wire.Struct(new(Repository), "DB"),
	)
