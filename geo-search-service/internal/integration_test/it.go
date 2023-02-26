//go:build integration
// +build integration

package integrationtest_test

import (
	di "geo-search/internal/integration_test/di_test"
	"geo-search/internal/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestIntegration(t *testing.T) {
	fixtures := di.InitializeFixtures()
	if err := fixtures.InsertFixtures(); err != nil {
		utils.LogFatalf("cannot insert fixtures: %v", err)
	}

	repoSuite := di.InitializeRepositorySuite()
	suite.Run(t, repoSuite)
}
