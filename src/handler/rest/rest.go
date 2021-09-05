package rest

import (
	"github.com/trongtb88/merchantsvc/src/middleware"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/trongtb88/merchantsvc/src/business/usecase"
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

	rst.mux.HandleFunc("/v1/accounts", middleware.Authenticate(rst.CreateMerchantAccount)).Methods(http.MethodPost)
	rst.mux.HandleFunc("/v1/accounts", middleware.Authenticate(rst.GetMerchantAccounts)).Methods(http.MethodGet)
	rst.mux.HandleFunc("/v1/accounts", middleware.Authenticate(rst.UpdateMerchantAccounts)).Methods(http.MethodPut)
	rst.mux.HandleFunc("/v1/accounts/{account_id}", middleware.Authenticate(rst.DeleteMerchantAccounts)).Methods(http.MethodDelete)

	rst.mux.HandleFunc("/v1/accounts/members", middleware.Authenticate(rst.CreateMerchantMember)).Methods(http.MethodPost)
	rst.mux.HandleFunc("/v1/accounts/members", middleware.Authenticate(rst.GetMerchantMembers)).Methods(http.MethodGet)
	rst.mux.HandleFunc("/v1/accounts/members", middleware.Authenticate(rst.UpdateMerchantMember)).Methods(http.MethodPut)
	rst.mux.HandleFunc("/v1/accounts/members/{member_id}", middleware.Authenticate(rst.DeleteMerchantMember)).Methods(http.MethodDelete)

	rst.mux.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

