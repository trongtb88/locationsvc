package location

import (
	"context"
	"fmt"
	"github.com/trongtb88/locationsvc/src/business/entity"
	"googlemaps.github.io/maps"
	"log"
)

func (loc *location) GetLocationsNearBy(ctx context.Context, param entity.LocationNearByParams) ([]entity.Location, entity.Pagination,error) {
	var (
		locs []entity.Location
		pagination entity.Pagination
		err error
	)

	latLngLocation, err := loc.geocodeAddress(param.StreetName)
	if err != nil {
		return locs,  pagination, err
	}

	placeType, err := loc.parseLocationType(param.PlaceType)

	r := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: latLngLocation.Lat,
			Lng: latLngLocation.Lng,
		},
		Radius:    param.Radius,
		Language:  "en",
		Type: placeType,
		PageToken: param.PageToken,
	}

	placesSearchResponse, err := loc.mapClient.NearbySearch(context.Background(), r)
	if err != nil {
		return locs, pagination, err
	}
	for _, result  := range placesSearchResponse.Results {
		fmt.Println(result.Name)
	}
	log.Println(placesSearchResponse)


	return locs, pagination, err
}

func (loc *location) parseLocationType (paramLocType string ) (maps.PlaceType, error){
	return maps.ParsePlaceType(paramLocType)
}

func (loc *location) geocodeAddress(address string) (entity.LatLng, error) {

	var latLngLocation entity.LatLng

	r := &maps.GeocodingRequest{
		Address: address,
	}

	res, err := loc.mapClient.Geocode(context.Background(), r)
	if err != nil || len(res) == 0 {
		return latLngLocation, fmt.Errorf("res Geocode err: %v", err)
	}

	latLngLocation.Lat = res[0].Geometry.Location.Lat
	latLngLocation.Lng = res[0].Geometry.Location.Lng
	return latLngLocation, nil
}
