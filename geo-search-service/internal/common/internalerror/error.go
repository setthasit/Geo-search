package internalerror

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
)

var NoInternalError = errors.New("no internal error")

type ApiError struct {
	Message         string
	HttpCode        int
	GrpcCode        codes.Code
	InternalCode    string
	InternalMessage string
}

var _ error = &ApiError{}

func NewApiError(httpCode int, grpcCode codes.Code, msg, internalCode, internalMsg string) *ApiError {
	return &ApiError{
		Message:         msg,
		HttpCode:        httpCode,
		GrpcCode:        grpcCode,
		InternalCode:    internalCode,
		InternalMessage: internalMsg,
	}
}

func (err *ApiError) Error() string {
	return fmt.Sprintf(
		"message: %s\n "+
			"http code: %d\n"+
			"gRpc code: %d\n"+
			"internal code: %s\n"+
			"internal message: %s\n",
		err.Message, err.HttpCode, err.GrpcCode, err.InternalCode, err.InternalMessage)
}
