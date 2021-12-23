package location

import (
	"context"
	"errors"
	"fmt"
	"github.com/trongtb88/locationsvc/src/common"
	"log"
	"sync"
	"time"

	"github.com/trongtb88/locationsvc/src/business/entity"
	"googlemaps.github.io/maps"
)

func (loc *location) GetLocationsNearBy(ctx context.Context, param entity.LocationNearByParams) ([]entity.Location, entity.Pagination,error) {
	var (
		locs []entity.Location
		pagination entity.Pagination
		err error
	)

	latLngLocation, err := loc.geocodeAddress(param.StreetName)
	if err != nil {
		log.Printf("Error when get lat, lng for address %v", err)
		return locs,  pagination, err
	}

	placeType, err := loc.parseLocationType(param.PlaceType)

	if err != nil {
		return locs, pagination, err
	}

	r := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: latLngLocation.Lat,
			Lng: latLngLocation.Lng,
		},
		Radius:    param.Radius * 1000,
		Language:  common.EN_LANGUAGE,
		Type: placeType,
		PageToken: param.PageToken,
	}

	placesSearchResponse, err := loc.mapClient.NearbySearch(context.Background(), r)



	if err != nil {
		log.Printf("Error when search near by %v", err)
		return locs, pagination, err
	}

	wg := sync.WaitGroup{}
	placesChannel := make(chan entity.Place, len(placesSearchResponse.Results))

	for _, result  := range placesSearchResponse.Results {
		wg.Add(1)
		go func(g *sync.WaitGroup, ctx context.Context, placeId string) {
			defer g.Done()
			log.Printf("Get address for place_id %s", placeId)
			addressLoc, err := loc.getAddress(ctx, placeId)
			if err != nil {
				log.Printf("Error when get place for place_id %s %s", placeId, err)
				placesChannel <- entity.Place{
					PlaceId: placeId,
					FormattedAddress: "",
				}
			}
			placesChannel <- addressLoc
		}(&wg, ctx, result.PlaceID)
	}

	wg.Wait()
	close(placesChannel)

	for place := range placesChannel {
		locs = append(locs, entity.Location{
			Name: place.Name,
			Address:  place.FormattedAddress,
		})
	}

	pagination.HasNextPage = true
	if placesSearchResponse.NextPageToken == "" {
		pagination.HasNextPage = false
	}
	pagination.NextPageToken = placesSearchResponse.NextPageToken
	return locs, pagination, err
}

func (loc *location) parseLocationType (paramLocType string ) (maps.PlaceType, error){
	return maps.ParsePlaceType(paramLocType)
}

func (loc *location) getAddress(ctx context.Context, placeId string) (entity.Place, error) {

	place, err := loc.SQLGetPlaceById(placeId)
	if err != nil  {
		log.Printf("Error on get place %v", err)
	}
	if len(place.PlaceId) > 0 {
		return place, nil
	}

	r := &maps.PlaceDetailsRequest{
		PlaceID: placeId,
	}
	// We can use redis/mysql to cache here to reduce API call to get addresses for placeIds faster
	// This version we use mysql to cache with PK = placeId
	placeDetail, err := loc.mapClient.PlaceDetails(ctx, r)
	if err != nil {
		return entity.Place{}, err
	}
	place = entity.Place{
		PlaceId:          placeDetail.PlaceID,
		Name:             placeDetail.Name,
		FormattedAddress: placeDetail.FormattedAddress,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}
	place, err = loc.SQLUpsertPlace(place)
	if err != nil {
		log.Printf("Error on save place %s", err)
	}
	return place, nil
}

func (loc *location) geocodeAddress(address string) (entity.LatLng, error) {
	var latLngLocation entity.LatLng
	r := &maps.GeocodingRequest{
		Address: address,
	}

	res, err := loc.mapClient.Geocode(context.Background(), r)
	if err != nil {
		return latLngLocation, fmt.Errorf("res Geocode err: %v", err)
	}
	if len(res) == 0 {
		return latLngLocation, errors.New("Not found lat, lng for input location")
	}

	// Only get the first match
	latLngLocation.Lat = res[0].Geometry.Location.Lat
	latLngLocation.Lng = res[0].Geometry.Location.Lng
	log.Printf("Lat, Lnf of %s is %d %d", address, latLngLocation.Lat, latLngLocation.Lng)

	return latLngLocation, nil
}
