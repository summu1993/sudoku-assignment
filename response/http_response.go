package response

import (
	"encoding/json"
)

type HttpResponse struct {
	RequestId     string `json:"request_id"`
	ErrorResponse struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		SubCode int    `json:"sub_code"`
	} `json:"error"`
	Data json.RawMessage `json:"data"`
}
