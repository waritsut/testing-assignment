package models

import "time"

type CashDrawer struct {
	Id          uint    `gorm:"primaryKey"`
	Money_Value float64 `gorm:"primaryKey,not null"`
	Amount      uint    `gorm:"not null"`
	Created_At  time.Time
	Updated_At  time.Time
}

func (CashDrawer) TableName() string {
	return "cash_drawers"
}
