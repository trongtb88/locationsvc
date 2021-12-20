package rest

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"


	"github.com/trongtb88/locationsvc/src/business/entity"
)

var decoderMember = schema.NewDecoder()


// Find name and address of 1 type of places (restaurants) located within a N kilometer radius  around 1 specific street name
// @Summary Find name and address of 1 type of places (restaurants) located within a N kilometer radius  around 1 specific street name
// @Description Find name and address of 1 type of places (restaurants) located within a N kilometer radius  around 1 specific street name
// @Tags NearByLocations
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param street_name query string " " "Street Name"
// @Param place_type query string restaurant "Number of km around "
// @Param radius query integer 2 "Number of km around "
// @Param page query int false " "
// @Param limit query int false " "
// @Success 200 {object} rest.ResponseGetAccounts
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/locations/nearby [get]
func (rst *rest) GetLocationsNearBy(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "GetLocationsNearByError",
			Message: err.Error(),
		})
		return
	}

	var param entity.LocationNearByParams
	fmt.Println(r.Form)

	err = decoderMember.Decode(&param, r.Form)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "GetLocationsNearByError",
			Message: err.Error(),
		})
		return
	}

	locations, pagination, err := rst.uc.Location.GetLocationsNearBy(r.Context(), param)

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "GetLocationsNearByError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, locations, &pagination)
}
