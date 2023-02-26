package repositories

import (
	"context"
	"geo-search/internal/entities"
)

//go:generate mockgen -source estate.go -destination mock_repositories/mock_estate.go -package mock_repositories

type EstateRepository interface {
	GetAll(ctx context.Context) ([]entities.Estate, error)
	GetByID(ctx context.Context, id uint) (*entities.Estate, error)
	GetByBoundBox(ctx context.Context, topLeft *entities.Location, bottomRight *entities.Location) ([]entities.Estate, int, error)

	Create(ctx context.Context, estate *entities.Estate) error
}
