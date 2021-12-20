package location

import (
	"context"
	Location "github.com/trongtb88/locationsvc/src/business/domain/location"
	"github.com/trongtb88/locationsvc/src/business/entity"
)

// Usecaseitf uc interface
type Usecaseitf interface {
	GetLocationsNearBy(ctx context.Context, param entity.LocationNearByParams) ([]entity.Location, entity.Pagination,error)
}

type location struct {
	location Location.DomainItf
}

// Options for uc, that is for config, can put them into consul or mapping with secret manager
type Options struct {
}

// InitLocation
func InitLocation(
	Location Location.DomainItf,
) Usecaseitf {
	return &location{
		Location,
	}
}
