package weather

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/handler"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/services"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	services.NewWeatherService,
	handler.NewWeatherHandler,
)
