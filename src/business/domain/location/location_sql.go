package location

import (
	"github.com/trongtb88/locationsvc/src/business/entity"
	"log"
)

func (loc *location) SQLGetPlaceById(placeId string) (entity.Place, error) {
	var (
		result entity.Place
	)

	place := entity.Place{PlaceId : placeId}
	tx := loc.sql.Model(&entity.Place{}).Where(&place).Find(&result)

	if tx.Error != nil {
		log.Printf("Error on tx", tx)
		return result, tx.Error
	}

	return result, nil
}

func (loc *location) SQLUpsertPlace(place entity.Place) (entity.Place, error) {
	tx := loc.sql.Save(&place)
	if tx.Error != nil {
		return place, tx.Error
	}
	return place, nil
}
