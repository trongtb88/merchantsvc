package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/trongtb88/merchantsvc/docs"
	resthandler "github.com/trongtb88/merchantsvc/src/handler/rest"
	// Business Layer Dep
	domain "github.com/trongtb88/merchantsvc/src/business/domain"
	usecase "github.com/trongtb88/merchantsvc/src/business/usecase"
	"github.com/trongtb88/merchantsvc/src/cmd/db"
)

var (
	sqlClient0 *gorm.DB

	// Server Infrastructure

	// Business Layer
	dom *domain.Domain
	uc  *usecase.Usecase
)

// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
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

	db := db.ConnectDB(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	// Business layer Initialization
	dom = domain.Init(
		db,
	)
	uc = usecase.Init(dom)

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8089"
	}

	router := mux.NewRouter()

	docs.SwaggerInfo.Title = os.Getenv("Meta_Namespace")
	docs.SwaggerInfo.Description = os.Getenv("Meta_Description")
	docs.SwaggerInfo.Version = os.Getenv("Meta_Version")
	docs.SwaggerInfo.BasePath = os.Getenv("Meta_BasePath")
	docs.SwaggerInfo.Host = os.Getenv("Meta_Host")

	// REST Handler Initialization
	_ = resthandler.Init(router, uc)

	log.Println("Starting server at port: ", serverPort)

	err = http.ListenAndServe(":"+serverPort, router)
	if err != nil {
		log.Println(err)
	}

}
