package http

import (
	"geo-search/internal/common/internalerror"

	"google.golang.org/grpc/codes"
)

func checkError(err error) (codes.Code, string) {
	intErr, ok := err.(*internalerror.ApiError)
	if ok {
		return intErr.GrpcCode, intErr.Message
	}

	return codes.Internal, "internal server error"
}
