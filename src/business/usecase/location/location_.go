package location

import (
	"context"
	"github.com/trongtb88/locationsvc/src/business/entity"
)

func (loc *location) GetLocationsNearBy(ctx context.Context, param entity.LocationNearByParams) ([]entity.Location, entity.Pagination, error) {
	// Can do some business logic in usecase, call cross domain here
	return loc.location.GetLocationsNearBy(ctx, param)
}
