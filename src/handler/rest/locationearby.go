package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Find name and address of 1 type of places (restaurants) located within a N kilometer radius  around 1 specific street name
// @Summary Find name and address of 1 type of places (restaurants) located within a N kilometer radius  around 1 specific street name
// @Description Find name and address of 1 type of places (restaurants) located within a N kilometer radius  around 1 specific street name
// @Tags NearByLocations
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param radius query integer 2 "Number of km around "
// @Param street_name query string " " "Street Name"
// @Param page query int false " "
// @Param limit query int false " "
// @Success 200 {object} rest.ResponseGetAccounts
// @Failure 400 {object} rest.HTTPErrResp
// @Failure 401 {object} rest.HTTPErrResp
// @Failure 500 {object} rest.HTTPErrResp
// @Router /v1/locations/nearby [get]
func (rst *rest) GetLocationsNearBy(w http.ResponseWriter, r *http.Request) {
	var param entity.CreateAccountParam
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:     "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	if err := json.Unmarshal(body, &param); err != nil {
		rst.httpRespError(w, r, http.StatusBadRequest, entity.ErrorMessage{
			Code:     "InvalidJsonBodyRequest",
			Message: err.Error(),
		})
		return
	}

	accounts, _,  err := rst.uc.Account.GetAccountsByParam(r.Context(), entity.GetAccountParam{
		Code:  param.Code,
	})

	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:     "CreateAccountError",
			Message: err.Error(),
		})
		return
	}

	if len(accounts) > 0 {
		rst.httpRespError(w, r,http.StatusBadRequest, entity.ErrorMessage{
			Code:     "DuplicateMerchantCode",
			Message: "Please choose other merchant code",
		})
		return
	}

	account, err := rst.uc.Account.CreateAccount(r.Context(), param)
	if err != nil {
		rst.httpRespError(w, r, http.StatusInternalServerError, entity.ErrorMessage{
			Code:     "CreateAccountError",
			Message: err.Error(),
		})
		return
	}

	rst.httpRespSuccess(w, r, http.StatusCreated, account, nil)
}
