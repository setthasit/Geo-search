//go:build integration
// +build integration

package fixtures

import "context"

type Fixtures struct {
	estateFixtures *EstateFixtures
}

func WireFixtures(estateFixtures *EstateFixtures) *Fixtures {
	return &Fixtures{
		estateFixtures: estateFixtures,
	}
}

func (fixture *Fixtures) InsertFixtures() error {
	ctx := context.Background()

	err := fixture.estateFixtures.InsertFixtures(ctx)
	if err != nil {
		return err
	}

	return nil
}
