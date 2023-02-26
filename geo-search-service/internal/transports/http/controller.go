package http

import "github.com/labstack/echo/v4"

type BaseController interface {
	Register(app *echo.Group)
}

type ControllerResponse struct {
	Code         int         `json:"code"`
	Data         interface{} `json:"data,omitempty"`
	Pagination   *Pagination `json:"pagination,omitempty"`
	Message      string      `json:"message,omitempty"`
	ErrorCode    string      `json:"error_code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

type Pagination struct {
	Page  int `query:"page" json:"page"`
	Size  int `query:"size" json:"size"`
	Count int `query:"count" json:"count"`
}
