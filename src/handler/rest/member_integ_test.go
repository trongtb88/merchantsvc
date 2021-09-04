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

func TestCreateSuccessCreateMerchantMember(t *testing.T) {
	randomEmail := uuid.New().String() + "@gmail.com"
	member := entity.CreateMemberParam{
		Email: randomEmail,
		MerchantId: 1,
		Name: "Test create email",
	}
	body, _ := json.Marshal(member)
	req, err := http.NewRequest("POST", "/v1/accounts/members", bytes.NewReader(body))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(e.CreateMerchantMember)
	handler.ServeHTTP(rr, req)

	var response ResponseCreateMember

	err = json.Unmarshal([]byte(rr.Body.String()), &response)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, randomEmail, response.Data.MemberEmail)
}

func TestGetMembersByParams(t *testing.T) {
	url := "/v1/accounts/members"

	randomEmail := uuid.New().String() + "@gmail.com"
	member := entity.CreateMemberParam{
		Email: randomEmail,
		MerchantId: 1,
		Name: "Test create email",
	}
	body, _ := json.Marshal(member)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(e.CreateMerchantMember)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)


	//-------------------Test Get By Email----------------------
	var response ResponseGetMembers
	req2, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	req2.URL.RawQuery = "email="+randomEmail
	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(e.GetMerchantMembers)
	handler2.ServeHTTP(rr2, req2)


	err = json.Unmarshal([]byte(rr2.Body.String()), &response)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr2.Code, http.StatusOK)
	assert.Equal(t, response.Data[0].MemberEmail, randomEmail)
	assert.Equal(t, response.Pagination.TotalElements > 0, true)
}

func TestUpdateMember(t *testing.T) {
	url := "/v1/accounts/members"
	randomEmail := uuid.New().String() + "@gmail.com"
	member := entity.CreateMemberParam{
		Email: randomEmail,
		MerchantId: 1,
		Name: "Test create email",
	}
	body, _ := json.Marshal(member)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(e.CreateMerchantMember)
	handler.ServeHTTP(rr, req)


	var responseCreated ResponseCreateMember
	err = json.Unmarshal([]byte(rr.Body.String()), &responseCreated)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, http.StatusCreated)

	//-------------------Test Update Member -----------------------
	memberAcc := entity.UpdateMerchantMember{
		Id: responseCreated.Data.Id,
		Email: responseCreated.Data.MemberEmail,
		Name: "New Name Member",
	}

	bodyUpdate, _ := json.Marshal(memberAcc)
	req2, err := http.NewRequest("PUT", url, bytes.NewReader(bodyUpdate))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(e.UpdateMerchantMember)
	handler2.ServeHTTP(rr2, req2)

	var responseUpdated ResponseCreateMember

	err = json.Unmarshal([]byte(rr2.Body.String()), &responseUpdated)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr2.Code, http.StatusOK)
	assert.Equal(t, responseUpdated.Data.MemberName, "New Name Member")
}

func TestDeleteMember(t *testing.T) {
	url := "/v1/accounts/members"
	randomEmail := uuid.New().String() + "@gmail.com"
	member := entity.CreateMemberParam{
		Email: randomEmail,
		MerchantId: 1,
		Name: "Test create email",
	}
	body, _ := json.Marshal(member)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(e.CreateMerchantMember)
	handler.ServeHTTP(rr, req)


	var responseCreated ResponseCreateMember
	err = json.Unmarshal([]byte(rr.Body.String()), &responseCreated)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, http.StatusCreated)

	//-------------------Test Delete Merchant Member -----------------------
	url = "/v1/accounts/members/{member_id}"
	replacer := strings.NewReplacer("{member_id}", strconv.FormatInt(responseCreated.Data.Id, 10))
	url = replacer.Replace(url)

	req2, err := http.NewRequest("DELETE", url, nil)
	req2 = mux.SetURLVars(req2, map[string]string{"member_id": strconv.FormatInt(responseCreated.Data.Id, 10)})

	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(e.DeleteMerchantMember)
	handler2.ServeHTTP(rr2, req2)

	assert.Equal(t, http.StatusOK, rr2.Code)
}
