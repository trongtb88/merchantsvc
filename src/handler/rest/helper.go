package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/trongtb88/merchantsvc/src/business/entity"
	"github.com/trongtb88/merchantsvc/src/common"
)

func (e *rest) httpRespSuccess(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}, p *entity.Pagination) {

	meta := Meta{
		Path:       r.URL.String(),
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Error:      ErrorMessage{
			Code:    "OK",
			Message: "Success",
		},
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	var (
		resp        interface{}
		err        error
	)

	if p != nil {
		resp = &ResponseSuccessPagination{
			Meta: meta,
			Data: data,
			Pagination: *p,
		}
	} else {
		resp = &ResponseSuccessNonPagination{
			Meta: meta,
			Data: data,
		}
	}

	raw, err := e.Marshal(resp)


	if err != nil {
		e.httpRespError(w, r, http.StatusInternalServerError, ErrorMessage{
			Code:    "StatusInternalServerError",
			Message: err.Error(),
		})
		return
	}

	w.Header().Set(common.HttpHeaderContentType, common.HttpContentJSON)
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func (e *rest) Marshal(resp interface{}) ([]byte, error) {
	return json.Marshal(&resp)
}

//func (rst *rest) httpRespError(w http.ResponseWriter, statusCode int, message ErrorMessage) {
//	raw, _ := json.Marshal(message)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(statusCode)
//	_, _ = w.Write(raw)
//}

func (e *rest) httpRespError(w http.ResponseWriter, r *http.Request, statusCode int, errorMsg ErrorMessage) {

	jsonErrResp := &HTTPErrResp{
		Meta: Meta{
			Path:       r.URL.String(),
			StatusCode: statusCode,
			Status:     http.StatusText(statusCode),
			Error:      errorMsg,
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}

	raw, err := e.Marshal(jsonErrResp)
	if err != nil {
		statusCode = http.StatusInternalServerError
		log.Println(err)
	}

	w.Header().Set(common.HttpHeaderContentType, common.HttpContentJSON)
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}
