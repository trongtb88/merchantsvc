package rest

import (
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

	rst.mux.HandleFunc("/v1/accounts", rst.CreateMerchantAccount).Methods(http.MethodPost)
	rst.mux.HandleFunc("/v1/accounts", rst.GetMerchantAccounts).Methods(http.MethodGet)
	rst.mux.HandleFunc("/v1/accounts", rst.UpdateMerchantAccounts).Methods(http.MethodPut)
	rst.mux.HandleFunc("/v1/accounts/{account_id}", rst.DeleteMerchantAccounts).Methods(http.MethodDelete)

	rst.mux.HandleFunc("/v1/accounts/members", rst.CreateMerchantMember).Methods(http.MethodPost)
	rst.mux.HandleFunc("/v1/accounts/members", rst.GetMerchantMembers).Methods(http.MethodGet)
	rst.mux.HandleFunc("/v1/accounts/members", rst.UpdateMerchantAccounts).Methods(http.MethodPut)
	rst.mux.HandleFunc("/v1/accounts/members/{member_id}", rst.DeleteMerchantAccounts).Methods(http.MethodDelete)

	rst.mux.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

