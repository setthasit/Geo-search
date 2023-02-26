package http

import (
	"geo-search/internal/entities"
	"geo-search/internal/interactors"
	"geo-search/internal/transports/http/model"
	"geo-search/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EstateController struct {
	estateInteractor *interactors.EstateInteractor
}

var _ BaseController = &EstateController{}

func WireEstateController(estateInteractor *interactors.EstateInteractor) *EstateController {
	return &EstateController{
		estateInteractor: estateInteractor,
	}
}

func (aCtrl *EstateController) Register(app *echo.Group) {
	api := app.Group("/estate")

	api.GET("", aCtrl.FindByBoxBound)
	api.POST("", aCtrl.Create)
}

func (aCtrl *EstateController) FindByBoxBound(c echo.Context) error {
	var request model.GetEstatesByBoundBoxRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
		})
	}

	err = c.Validate(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerResponse{
			Code:    http.StatusBadRequest,
			Message: utils.ApiValidateError(err).Error(),
		})
	}

	ctx := c.Request().Context()
	estates, count, err := aCtrl.estateInteractor.FindByBoundBox(
		ctx,
		&entities.Location{Lat: request.TopLeftLat, Lon: request.TopLeftLon},
		&entities.Location{Lat: request.BottomRightLat, Lon: request.BottomRightLon},
	)
	if err != nil {
		return c.JSON(checkError(err))
	}

	return c.JSON(http.StatusOK, ControllerResponse{
		Code: http.StatusOK,
		Data: estates,
		Pagination: &Pagination{
			Page:  1,
			Size:  1000,
			Count: count,
		},
	})
}

func (aCtrl *EstateController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var request model.CreateEstateRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
		})
	}

	err = c.Validate(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerResponse{
			Code:    http.StatusBadRequest,
			Message: utils.ApiValidateError(err).Error(),
		})
	}

	newEstate := model.MapCreateEstateRequestToEntity(&request)
	estate, err := aCtrl.estateInteractor.Create(ctx, newEstate)
	if err != nil {
		return c.JSON(checkError(err))
	}

	return c.JSON(http.StatusOK, ControllerResponse{
		Code: http.StatusOK,
		Data: estate,
	})
}
