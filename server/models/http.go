package models

import "net/http"

type HttpResponse struct {
	Data   interface{} `json:"data"`
	Status HttpStatus  `json:"status"`
}

type HttpStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HttpResponseStatusOk() *HttpResponse {
	return &HttpResponse{
		Status: HttpStatus{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}
}
