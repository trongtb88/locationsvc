package domain

import (
	"github.com/trongtb88/locationsvc/src/business/domain/location"
	"googlemaps.github.io/maps"
	"gorm.io/gorm"
)

type Domain struct {
	Location       location.DomainItf
}

func Init(
	sql *gorm.DB,
	mapClient *maps.Client,
) *Domain {

	return &Domain{
		Location: location.InitLocationDomain(
			sql,
			mapClient,
		),
	}
}

