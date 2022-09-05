package seed

import (
	"cashier-service/internal/migrations"
	"cashier-service/internal/pkg/db_driver"
	"log"
)

func Load(db db_driver.Repo) {
	DB := db.GetDb()

	if err := DB.Exec("DROP TABLE IF EXISTS cash_drawers, migrations;").Error; err != nil {
		log.Panicf("Could not drop table: %v", err)
	}

	migrations.Migrate(db)

	if err := DB.Exec(`INSERT INTO "cash_drawers"("money_value","amount","created_at","updated_at")
	VALUES
	(1000, 10, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(500,  20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'), 
	(100,  15, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(50,	 20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(20,	 30, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(10,	 20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(5, 	 20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(1,		 20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(0.25, 50, '2022-01-01 00:00:00', '2022-01-01 00:00:00');`).Error; err != nil {
		log.Panicf("Could not insert data: %v", err)
	}

	log.Println("Creating cash_drawers data")
}
