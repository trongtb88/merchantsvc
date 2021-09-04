package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/trongtb88/merchantsvc/src/business/entity"
)

func TestCreateSuccessCreateMerchantAccount(t *testing.T) {
	code := uuid.New().String()
	acc := entity.CreateAccountParam{
		Code: code,
		Name: "Test create account",
	}
	body, _ := json.Marshal(acc)
	req, err := http.NewRequest("POST", "/v1/accounts", bytes.NewReader(body))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(e.CreateMerchantAccount)
	handler.ServeHTTP(rr, req)

	var response ResponseCreateAccount

	err = json.Unmarshal([]byte(rr.Body.String()), &response)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, http.StatusCreated)
	assert.Equal(t, response.Data.MerchantCode, code)
}

func TestGetAccountsByParams(t *testing.T) {
	url := "/v1/accounts"

	code := uuid.New().String()
	acc := entity.CreateAccountParam{
		Code: code,
		Name: "Test create account",
	}
	body, _ := json.Marshal(acc)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(e.CreateMerchantAccount)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)


	//-------------------Test Get By Code-----------------------
	var response ResponseGetAccounts
	req2, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	req2.URL.RawQuery = "code="+code
	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(e.GetMerchantAccounts)
	handler2.ServeHTTP(rr2, req2)


	err = json.Unmarshal([]byte(rr2.Body.String()), &response)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr2.Code, http.StatusOK)
	assert.Equal(t, response.Data[0].MerchantCode, code)
	assert.Equal(t, response.Pagination.TotalElements > 0, true)
}

func TestUpdateAccount(t *testing.T) {
	url := "/v1/accounts"

	code := uuid.New().String()
	acc := entity.CreateAccountParam{
		Code: code,
		Name: "Test create account",
	}
	body, _ := json.Marshal(acc)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(e.CreateMerchantAccount)
	handler.ServeHTTP(rr, req)

	var responseCreated ResponseCreateAccount

	err = json.Unmarshal([]byte(rr.Body.String()), &responseCreated)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusCreated)


	//-------------------Test Update Account -----------------------
	updatedAcc := entity.UpdateMerchantAccount{
		Id: responseCreated.Data.Id,
		MerchantCode: responseCreated.Data.MerchantCode,
		MerchantName: "New Name",
	}

	bodyUpdate, _ := json.Marshal(updatedAcc)
	req2, err := http.NewRequest("PUT", url, bytes.NewReader(bodyUpdate))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(e.UpdateMerchantAccounts)
	handler2.ServeHTTP(rr2, req2)

	var responseUpdated ResponseCreateAccount

	err = json.Unmarshal([]byte(rr2.Body.String()), &responseUpdated)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr2.Code, http.StatusOK)
	assert.Equal(t, responseUpdated.Data.MerchantName, "New Name")
}

func TestDeleteAccount(t *testing.T) {
	url := "/v1/accounts"
	code := uuid.New().String()
	acc := entity.CreateAccountParam{
		Code: code,
		Name: "Test create account",
	}
	body, _ := json.Marshal(acc)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(e.CreateMerchantAccount)
	handler.ServeHTTP(rr, req)

	var responseCreated ResponseCreateAccount

	err = json.Unmarshal([]byte(rr.Body.String()), &responseCreated)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusCreated)

	//-------------------Test Delete Merchant Account -----------------------
	url = "/v1/accounts/{account_id}"
	replacer := strings.NewReplacer("{account_id}", strconv.FormatInt(responseCreated.Data.Id, 10))
	url = replacer.Replace(url)

	req2, err := http.NewRequest("DELETE", url, nil)
	req2 = mux.SetURLVars(req2, map[string]string{"account_id": strconv.FormatInt(responseCreated.Data.Id, 10)})

	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(e.DeleteMerchantAccounts)
	handler2.ServeHTTP(rr2, req2)

	assert.Equal(t, http.StatusOK, rr2.Code)
}
