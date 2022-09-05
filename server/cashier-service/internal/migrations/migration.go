package migrations

import (
	"cashier-service/internal/pkg/db_driver"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Migrate(db db_driver.Repo) {
	DB := db.GetDb()
	m := gormigrate.New(
		DB,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1662215040CreateCashDrawerTable(),
		},
	)

	if err := m.Migrate(); err != nil {
		log.Panicf("Could not migrate: %v", err)
	}
}
