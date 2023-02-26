package http

import (
	"geo-search/internal/common/internalerror"
	"net/http"
)

func checkError(err error) (int, *ControllerResponse) {
	intErr, ok := err.(*internalerror.ApiError)
	if ok {
		return intErr.HttpCode, &ControllerResponse{
			Code:         intErr.HttpCode,
			Message:      intErr.Message,
			ErrorMessage: intErr.InternalMessage,
			ErrorCode:    intErr.InternalCode,
		}
	}

	return http.StatusInternalServerError, &ControllerResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}
}
