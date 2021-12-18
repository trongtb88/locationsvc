package rest

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/trongtb88/locationsvc/src/business/usecase"
	"github.com/trongtb88/locationsvc/src/middleware"
)

// REST rest interface
type REST interface{}

var once = &sync.Once{}

type rest struct {
	mux    *mux.Router
	uc     *usecase.Usecase
}

func Init(router *mux.Router, uc *usecase.Usecase) REST {
	var e *rest
	once.Do(func() {
		e = &rest{
			mux:    router,
			uc:     uc,
		}
		e.Serve()
	})
	return e
}


func (rst *rest) Serve() {
	rst.mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})

	rst.mux.HandleFunc("/v1/locations/nearby", middleware.Authenticate(rst.GetLocationsNearBy)).Methods(http.MethodGet)

	rst.mux.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

