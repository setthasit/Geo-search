package internalerror

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	ESTATE_INTERNAL_PREFIX = "EST"
)

// Interactor

// AdminInteractorInvalidUsernameOrPassword - Create internal error for invalid username/password
func EstateInteractor(err error) *ApiError {
	return NewApiError(
		http.StatusBadRequest,
		codes.InvalidArgument,
		"invalid username/password",
		fmt.Sprintf("%sINT_4000001", ESTATE_INTERNAL_PREFIX),
		err.Error(),
	)
}

// Repository

// EstateRepositoryIndexing - Create internal error for indexing error
func EstateRepositoryIndexing(err error) *ApiError {
	return NewApiError(
		http.StatusInternalServerError,
		codes.Internal,
		"cannot indexing document",
		fmt.Sprintf("%sREP_5000001", ESTATE_INTERNAL_PREFIX),
		err.Error(),
	)
}

// EstateRepositoryIndexing - Create internal error for indexing error
func EstateRepositoryParsingRequest(err error) *ApiError {
	return NewApiError(
		http.StatusInternalServerError,
		codes.Internal,
		"cannot parse request",
		fmt.Sprintf("%sREP_5000002", ESTATE_INTERNAL_PREFIX),
		err.Error(),
	)
}

// EstateRepositoryIndexing - Create internal error for indexing error
func EstateRepositoryParsingResponse(err error) *ApiError {
	return NewApiError(
		http.StatusInternalServerError,
		codes.Internal,
		"cannot parse response",
		fmt.Sprintf("%sREP_5000003", ESTATE_INTERNAL_PREFIX),
		err.Error(),
	)
}

// EstateRepositoryIndexing - Create internal error for indexing error
func EstateRepositoryRequest(err error) *ApiError {
	return NewApiError(
		http.StatusInternalServerError,
		codes.Internal,
		"cannot process request",
		fmt.Sprintf("%sREP_5000003", ESTATE_INTERNAL_PREFIX),
		err.Error(),
	)
}

// EstateRepositoryIndexing - Create internal error for indexing error
func EstateRepositoryCannotSave(err error) *ApiError {
	return NewApiError(
		http.StatusInternalServerError,
		codes.Internal,
		"cannot save estate to database",
		fmt.Sprintf("%sREP_5000004", ESTATE_INTERNAL_PREFIX),
		err.Error(),
	)
}
