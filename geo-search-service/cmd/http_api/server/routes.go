package server

import (
	"geo-search/internal/transports/http"

	"github.com/labstack/echo/v4"
)

type ApiRoutes struct {
	estateController *http.EstateController
}

func ProvideApiRoutes(
	estateController *http.EstateController,
) *ApiRoutes {
	return &ApiRoutes{
		estateController: estateController,
	}
}

func (route *ApiRoutes) Register(app *echo.Echo) {
	apiV1 := app.Group("/api/v1")

	route.estateController.Register(apiV1)
}
