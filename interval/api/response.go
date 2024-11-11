package api

import (
	"github.com/gin-gonic/gin"
)

const (
	ErrorCodeMissing       = "missing"
	ErrorCodeMissingField  = "missing_field"
	ErrorCodeInvalid       = "invalid"
	ErrorCodeAlreadyExists = "already_exists"
	ErrorCodeUnprocessable = "unprocessable"
	ErrorCodeCustom        = "custom"
)

// ErrorDetail represents individual error details
type ErrorDetail struct {
	Field   string `json:"field,omitempty"`   // The field that caused the error
	Message string `json:"message,omitempty"` // A detailed error message
	Code    string `json:"code,omitempty"`    // The error code
}

// SuccessResponse 用于封装 GitHub 风格的成功响应格式
type SuccessResponse struct {
	Data interface{} `json:"data,omitempty"`
}

// SuccessResponseHandler 用于返回成功的响应
func SuccessResponseHandler(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, SuccessResponse{
		Data: data,
	})
}

// ErrorResponse 用于封装错误响应格式
type ErrorResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	//DocumentationURL string      `json:"documentation_url"`
}

// ErrorResponseHandler is a function to return error responses in a standard format
func ErrorResponseHandler(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.JSON(statusCode, ErrorResponse{
		Message: message,
		Errors:  errors,
	})
}
