//go:build integration
// +build integration

package repositories

import (
	"context"
	"geo-search/internal/entities"
	"geo-search/internal/utils"
	"time"
)

func (s *RepositoryTestSuite) TestCreateEstate() {
	// Should avoid create estates in this boundary
	// topLeft := []float64{51.89563455208273, 4.49980791636062}
	// bottomRight := []float64{51.888317570989095, 4.511376402730734}
	s.Run("should create with given valid estate, Then return no error", func() {
		estate := &entities.Estate{
			Title: "test 1",
			Location: entities.Location{
				Lat: 10.00100,
				Lon: 80.00001,
			},
		}
		expectedEstate := &entities.Estate{
			Title: "ittest estate 1",
			Location: entities.Location{
				Lat: 10.00100,
				Lon: 80.00001,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		}

		err := s.EstateRepo.Create(context.TODO(), estate)

		s.Assert().NoError(err)
		s.Assert().Equal(expectedEstate.Title, estate.Title)
		s.Assert().Equal(expectedEstate.Location.Lat, estate.Location.Lat)
		s.Assert().Equal(expectedEstate.Location.Lon, estate.Location.Lon)
		s.Assert().True(utils.TimeSubMinutes(expectedEstate.CreatedAt, estate.CreatedAt) < 5)
		s.Assert().True(utils.TimeSubMinutes(expectedEstate.UpdatedAt, estate.UpdatedAt) < 5)
		s.Assert().Nil(estate.DeletedAt)
	})

	s.Run("should create with given minus {lat, long}, then return no error", func() {
		estate := &entities.Estate{
			Title: "test 1",
			Location: entities.Location{
				Lat: -0.12345,
				Lon: -100.67890,
			},
		}
		expectedEstate := &entities.Estate{
			Title: "ittest estate 1",
			Location: entities.Location{
				Lat: -0.12345,
				Lon: -100.67890,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		}

		err := s.EstateRepo.Create(context.TODO(), estate)

		s.Assert().NoError(err)
		s.Assert().Equal(expectedEstate.Title, estate.Title)
		s.Assert().Equal(expectedEstate.Location.Lat, estate.Location.Lat)
		s.Assert().Equal(expectedEstate.Location.Lon, estate.Location.Lon)
		s.Assert().True(utils.TimeSubMinutes(expectedEstate.CreatedAt, estate.CreatedAt) < 5)
		s.Assert().True(utils.TimeSubMinutes(expectedEstate.UpdatedAt, estate.UpdatedAt) < 5)
		s.Assert().Nil(estate.DeletedAt)
	})
}
