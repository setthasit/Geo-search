package model

import "geo-search/internal/entities"

type GetEstatesByBoundBoxRequest struct {
	TopLeftLat     float64 `query:"top_left_lat" validate:"required"`
	TopLeftLon     float64 `query:"top_left_lon" validate:"required"`
	BottomRightLat float64 `query:"bottom_right_lat" validate:"required"`
	BottomRightLon float64 `query:"bottom_right_lon" validate:"required"`
}

type CreateEstateRequest struct {
	Title    string   `json:"title" validate:"required"`
	Location Location `json:"location" validate:"required"`
}

type Location struct {
	Lat float64 `query:"lat" json:"lat" validate:"required"`
	Lon float64 `query:"lon" json:"lon" validate:"required"`
}

func MapCreateEstateRequestToEntity(request *CreateEstateRequest) *entities.Estate {
	return &entities.Estate{
		Title: request.Title,
		Location: entities.Location{
			Lat: request.Location.Lat,
			Lon: request.Location.Lon,
		},
	}
}
