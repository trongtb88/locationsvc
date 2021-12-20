package entity

type LocationNearByParams struct {
	StreetName string `schema:"street_name"`
	PlaceType string `schema:"place_type"`
	Radius uint `schema:"radius"`
	Limit int `schema:"limit"`
	Offset int `schema:"offset"`
	PageToken string `schema:"page_token"`
}

type Location struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

// LatLng represents a location on the Earth.
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}