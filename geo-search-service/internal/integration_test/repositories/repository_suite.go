//go:build integration
// +build integration

package repositories

import (
	"geo-search/internal/repositories"

	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite

	// Repositories
	EstateRepo repositories.EstateRepository
}

func WireRepositoryTestSuite(
	estateRepo repositories.EstateRepository,
) *RepositoryTestSuite {
	return &RepositoryTestSuite{
		EstateRepo: estateRepo,
	}
}
