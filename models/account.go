package models

import "time"

type Account struct {
	Id        int       `json:"id" gorm:"column:id"`
	Owner     string    `json:"owner" gorm:"column:owner"`
	Balance   float32   `json:"balance" gorm:"column:balance"`
	Currency  string    `json:"currency" gorm:"column:currency"`
	Create_at time.Time `json:"create_at" gorm:"create_at"`
}
