package rest

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/trongtb88/merchantsvc/src/business/domain"
	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/business/usecase"
	"github.com/trongtb88/merchantsvc/src/cmd/db"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var (
	sqlClient0     *gorm.DB

	// Server Infrastructure

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

	// Business layer Initialization
	dom = domain.Init(
		db,
	)
	uc = usecase.Init(dom)

	db.Debug().AutoMigrate(&entity.MerchantAccount{})
	db.Debug().AutoMigrate(&entity.MerchantMember{})

	router := mux.NewRouter()

	e = &rest{
		mux:    router,
		uc:     uc,
	}

	os.Exit(m.Run())

}
