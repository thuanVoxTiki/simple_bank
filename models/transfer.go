package models

import "time"

type Transfer struct {
	Id              int        `json:"id" gorm:"column:id"`
	From_account_id int        `json:"from_account_id" gorm:"column:from_account_id"`
	To_account_id   int        `json:"to_account_id" gorm:"column:to_account_id"`
	Amount          int        `json:"amount" gorm:"column:amount"`
	Create_at       *time.Time `json:"create_at" gorm:"column:create_at"`
}
