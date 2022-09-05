package cashier_service

func (s *CashierService) GetBalance() (balance float64, err error) {
	err = s.Raw(`SELECT SUM(money_value * amount) FROM cash_drawers;`).Scan(&balance).Error
	if err != nil {
		return balance, err
	}
	return balance, nil
}
