package rest

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLocationsNearBy(t *testing.T) {
	url := "/v1/locations/nearby"
	//-------------------Test Get Locations By Parameters----------------------
	var response ResponseGetLocationNearBy
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	streetName := "Sukhumvit, Thailand"
	placeType := "restaurant"
	radius := "2" // km
	req.URL.RawQuery = "street_name="+streetName + "&place_type="+placeType + "&radius="+radius
	rr := httptest.NewRecorder()
	handler2 := http.HandlerFunc(e.GetLocationsNearBy)
	handler2.ServeHTTP(rr, req)

	err = json.Unmarshal([]byte(rr.Body.String()), &response)
	if err != nil {
		log.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t,len(response.Data) > 0, true)
}