package server

import (
	"geo-search/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApiServer struct {
	app    *echo.Echo
	routes *ApiRoutes
}

func ProvideApiServer(routes *ApiRoutes) *ApiServer {
	app := echo.New()
	app.Use(
		middleware.CORS(),
		// middleware.CSRF(),
		middleware.Secure(),
		middleware.RequestID(),
		middleware.Logger(),
	)
	app.Validator = utils.NewAPIValidator(validator.New())

	return &ApiServer{
		app:    app,
		routes: routes,
	}
}

func (s *ApiServer) Run() {
	app := s.app

	s.routes.Register(app)

	app.Logger.Fatal(app.Start(":10000"))
}
