package rest

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/trongtb88/locationsvc/src/business/domain"
	"github.com/trongtb88/locationsvc/src/business/entity"
	"github.com/trongtb88/locationsvc/src/business/usecase"
	"github.com/trongtb88/locationsvc/src/cmd/db"
	"googlemaps.github.io/maps"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var (
	sqlClient0     *gorm.DB
	mapClient       *maps.Client

	// Server Infrastructure

	apiKey    = flag.String("key", "", "API Key for using Google Maps API.")
	clientID  = flag.String("client_id", "", "ClientID for Maps for Work API access.")
	signature = flag.String("signature", "", "Signature for Maps for Work API access.")

	// Business Layer
	dom *domain.Domain
	uc  *usecase.Usecase
	e   *rest
)

// We can improve integration tests by using csv files to make integration tests.
// 1 file for metadata, 1 file for req,1 file for response
// But in this scope, I will not use it.
func TestMain(m *testing.M) {

	var err error
	err = godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		log.Println("We are getting the env values")
	}

	db := db.ConnectDB (
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	*apiKey = os.Getenv("GOOGLE_MAP_API_KEY")

	//Init googlemap client
	if *apiKey != "" {
		mapClient, err = maps.NewClient(maps.WithAPIKey(*apiKey))
	} else if *clientID != "" || *signature != "" {
		mapClient, err = maps.NewClient(maps.WithClientIDAndSignature(*clientID, *signature))
	} else {
		log.Fatal("Please specify an API Key, or Client ID and Signature.")
	}

	// Business layer Initialization
	dom = domain.Init(
		db,
		mapClient,
	)
	uc = usecase.Init(dom)

	db.Debug().AutoMigrate(&entity.Place{})

	router := mux.NewRouter()

	e = &rest{
		mux:    router,
		uc:     uc,
	}

	os.Exit(m.Run())

}
