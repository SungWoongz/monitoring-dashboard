package models

import "net/http"

type HttpResponse struct {
	Data   interface{} `json:"data"`
	Status HttpStatus  `json:"status"`
}

type HttpStatus struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"OK"`
}

func HttpResponseStatusOk() *HttpResponse {
	return &HttpResponse{
		Status: HttpStatus{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}
}

func HttpResponseBadRequest(msg string) *HttpResponse {
	return &HttpResponse{
		Status: HttpStatus{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest) + msg,
		},
	}
}

func HttpResponseInternalServerError(msg string) *HttpResponse {
	return &HttpResponse{
		Status: HttpStatus{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError) + msg,
		},
	}
}
