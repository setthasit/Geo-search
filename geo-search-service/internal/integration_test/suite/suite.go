//go:build integration
// +build integration

package suite

// import (
// 	"geo-search/internal/integration_test/fixtures"
// 	"geo-search/internal/interactors"
// 	"geo-search/internal/repositories"
// 	"geo-search/internal/utils"

// 	testsuite "github.com/stretchr/testify/suite"
// )

// type IntegrationTestSuite struct {
// 	testsuite.Suite

// 	// Fixtures
// 	fixtures *fixtures.Fixtures

// 	// Interactor
// 	EstateIta *interactors.EstateInteractor

// 	// Repositories
// 	EstateRepo repositories.EstateRepository
// }

// func WireIntegrationTestSuite(
// 	// Fixtures
// 	fixtures *fixtures.Fixtures,
// 	// Interactors
// 	estateIta *interactors.EstateInteractor,
// 	// Repositories
// 	estateRepo repositories.EstateRepository,
// ) *IntegrationTestSuite {
// 	return &IntegrationTestSuite{
// 		// Fixtures
// 		fixtures: fixtures,
// 		// Interactors
// 		EstateIta: estateIta,
// 		// Repositories
// 		EstateRepo: estateRepo,
// 	}
// }

// func (it *IntegrationTestSuite) InsertFixtures() {
// 	err := it.fixtures.InsertFixtures()
// 	if err != nil {
// 		utils.LogFatalf("cannot insert fixtures: %v", err)
// 	}
// }
