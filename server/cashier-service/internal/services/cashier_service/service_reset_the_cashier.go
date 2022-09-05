package cashier_service

func (s *CashierService) ResetTheCashier() (err error) {
	err = s.Exec(`
	DELETE FROM cash_drawers;
	SELECT setval('cash_drawers_id_seq', 1);
	INSERT INTO "cash_drawers"("money_value","amount","created_at","updated_at")
	VALUES
	(1000, 10, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(500,  20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'), 
	(100,  15, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(50,	 20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(20,	 30, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(10,	 20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(5, 	 20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(1,		 20, '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
	(0.25, 50, '2022-01-01 00:00:00', '2022-01-01 00:00:00');`).Error
	if err != nil {
		return err
	}
	return nil
}