package rest

import (
	"github.com/gorilla/schema"
	"google.golang.org/appengine"
	"net/http"
	"strings"

	"github.com/trongtb88/locationsvc/src/business/entity"
)

var decoderMember = schema.NewDecoder()


// @Summary Find name and address of one kine of place (restaurants) located within a N kilometer radius  around 1 specific street name
// @Description Find name and address of 1 type of place (restaurants) located within a N kilometer radius  around 1 specific street name
// @Tags NearByLocations
// @Accept json
// @Produce json
// @Param street_name query string true "Street Name eg : Sukhumvit, Thailand"
// @Param place_type query string true "Place Type" Enums(restaurant, school)
// @Param radius query integer true "Radius in kilometer"
// @Param page_token query string false " "
// @Success 200 {object} rest.ResponseGetLocationNearBy
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


	err = decoderMember.Decode(&param, r.Form)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "GetLocationsNearByError",
			Message: err.Error(),
		})
		return
	}

	if len(strings.TrimSpace(param.StreetName)) == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "BadRequest",
			Message: "Street Name must be not empty",
		})
		return
	}

	if len(strings.TrimSpace(param.PlaceType)) == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "BadRequest",
			Message: "Place_type must be not empty",
		})
		return
	}

	if param.Radius == 0 {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:    "BadRequest",
			Message: "Radius must > 0",
		})
		return
	}

	locations, pagination, err := rst.uc.Location.GetLocationsNearBy(appengine.NewContext(r), param)

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:    "GetLocationsNearByError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusOK, locations, &pagination)
}