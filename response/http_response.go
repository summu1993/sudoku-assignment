package response

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
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

func Error(ctx *gin.Context, code int, message string, requestId string) {
	var response HttpResponse
	response.RequestId = requestId
	response.ErrorResponse.Message = message
	response.ErrorResponse.Code = code

	ctx.JSON(code, response)
}
