package location

import (
	"context"
	"github.com/trongtb88/locationsvc/src/business/entity"
	"googlemaps.github.io/maps"
	"gorm.io/gorm"
)

// DomainItf domain interface for location account
type DomainItf interface {
	GetLocationsNearBy(ctx context.Context, param entity.LocationNearByParams) ([]entity.Location, entity.Pagination,error)
}

type location struct {
	sql         *gorm.DB
	mapClient *maps.Client
}

// InitLocationDomain domain init
func InitLocationDomain(
	sql *gorm.DB,
	mapClient *maps.Client,
) DomainItf {
	loc := &location{
		sql:         sql,
		mapClient: mapClient,
	}
	return loc
}
