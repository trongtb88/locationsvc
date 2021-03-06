package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"

	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	resthandler "github.com/trongtb88/locationsvc/src/handler/rest"
	"github.com/trongtb88/locationsvc/docs"
	// Business Layer Dep
	domain "github.com/trongtb88/locationsvc/src/business/domain"
	usecase "github.com/trongtb88/locationsvc/src/business/usecase"
	"github.com/trongtb88/locationsvc/src/cmd/db"

)

var (
	sqlClient0     *gorm.DB
	mapClient       *maps.Client

	apiKey    = flag.String("key", "", "API Key for using Google Maps API.")
	clientID  = flag.String("client_id", "", "ClientID for Maps for Work API access.")
	signature = flag.String("signature", "", "Signature for Maps for Work API access.")

	// Server Infrastructure

	// Business Layer
	dom *domain.Domain
	uc  *usecase.Usecase
)

func main() {


	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// Init mysql
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

	serverPort := os.Getenv("SERVER_PORT")

	router := mux.NewRouter()

	docs.SwaggerInfo.Title = os.Getenv("Meta_Namespace")
	docs.SwaggerInfo.Description = os.Getenv("Meta_Description")
	docs.SwaggerInfo.Version = os.Getenv("Meta_Version")
	docs.SwaggerInfo.BasePath = os.Getenv("Meta_BasePath")
	docs.SwaggerInfo.Host = os.Getenv("Meta_Host")

	// REST Handler Initialization
	_ = resthandler.Init(router,  uc)

	log.Println("Starting server at port: ", serverPort)

	err = http.ListenAndServe(":"+serverPort, router)
	if err != nil {
		log.Println(err)
	}




}
