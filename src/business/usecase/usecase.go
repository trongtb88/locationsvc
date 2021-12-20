package usecase

import (
	"github.com/trongtb88/locationsvc/src/business/domain"
	"github.com/trongtb88/locationsvc/src/business/usecase/location"
)

type Usecase struct {
	Location location.Usecaseitf
}

// Init all usecase
func Init(
	dom *domain.Domain,
) *Usecase {

	return &Usecase{
		Location: location.InitLocation(
			dom.Location,
		),
	}
}

