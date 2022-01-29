package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
	Exception string `json:"exception"`
}

func SendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int, e error, success string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil && success == "" {
		return
	}

	if v, ok := data.(int); ok {
		data = PrepareErrorResponse(v, e.Error())
	} else {
		data = PrepareSuccessResponse(success, data)
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err:%v\n", err)
	}
}

func PrepareErrorResponse(code int, errMsg string) ErrorResponse {
	if errMsg == "sql: no rows in result set" {
		errMsg = "Record not found."
	}

	return ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			ErrorCode: code,
			Message:   errorsMap[code],
			Exception: errMsg,
		},
	}
}

func PrepareSuccessResponse(msg string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func Parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}
