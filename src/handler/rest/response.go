package rest

import "github.com/trongtb88/locationsvc/src/business/entity"

type ResponseSuccessNonPagination struct {
	Meta entity.Meta `json:"metadata"`
	Data interface{} `json:"data"`
}

type ResponseSuccessPagination struct {
	Meta       entity.Meta       `json:"metadata"`
	Data       interface{}       `json:"data"`
	Pagination entity.Pagination `json:"pagination"`
}

type ResponseGetLocationNearBy struct {
	Meta       entity.Meta       `json:"metadata"`
	Data       []entity.Location `json:"data"`
	Pagination entity.Pagination `json:"pagination"`
}

type HTTPErrResp struct {
	Meta entity.Meta `json:"metadata"`
}
type HTTPEmptyResp struct {
	Meta entity.Meta `json:"metadata"`
}
