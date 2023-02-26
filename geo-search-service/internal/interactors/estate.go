package interactors

import (
	"context"
	"geo-search/internal/entities"
	"geo-search/internal/repositories"
)

type EstateInteractor struct {
	estateRepo repositories.EstateRepository
}

func WireEstateInteractor(estateRepo repositories.EstateRepository) *EstateInteractor {
	return &EstateInteractor{
		estateRepo: estateRepo,
	}
}

func (iat *EstateInteractor) FindByBoundBox(ctx context.Context, topLeft *entities.Location, bottomRight *entities.Location) ([]entities.Estate, int, error) {
	estates, count, err := iat.estateRepo.GetByBoundBox(ctx, topLeft, bottomRight)
	if err != nil {
		return nil, 0, err
	}

	return estates, count, nil
}

func (iat *EstateInteractor) Create(ctx context.Context, estate *entities.Estate) (*entities.Estate, error) {
	err := iat.estateRepo.Create(ctx, estate)
	if err != nil {
		return nil, err
	}
	return estate, nil
}
