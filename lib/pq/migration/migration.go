package migration

import (
	"path/filepath"
	"rilutham/stovia/lib/utils"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	bindata "github.com/mattes/migrate/source/go-bindata"
)

// Direction nodoc
type Direction uint8

const (
	_ Direction = iota
	Up
	Down
	Drop
)

// Run :nodoc:
func Run(dsn string, direction Direction, step int) error {
	assetNamesWithoutPath := make([]string, len(utils.AssetNames()))

	for i, name := range utils.AssetNames() {
		_, name = filepath.Split(name)
		assetNamesWithoutPath[i] = name
	}

	s := bindata.Resource(assetNamesWithoutPath, func(name string) ([]byte, error) {
		return utils.Asset(filepath.Join("resources/sql/migration", name))
	})

	d, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		return err
	}

	step = intAbs(step)
	if step < 1 {
		step = 1
	}

	switch direction {
	case Up:
		return m.Up()
	case Down:
		return m.Steps(-step)
	case Drop:
		return m.Drop()
	}

	return nil
}

func intAbs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
