package migrations

import (
	"cashier-service/internal/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1662215040CreateCashDrawerTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1662215040",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.CashDrawer{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Exec("DROP TABLE cash_drawers;").Error
		},
	}

}
