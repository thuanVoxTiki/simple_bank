package models

import (
	"time"
)

type Entrie struct {
	Id        int        `json:"id" gorm:"id"`
	Acount_id int        `json:"acount_id" gorm:"acount_id"`
	Amount    float32    `json:"amount" gorm:"column:amount"`
	Create_at *time.Time `json:"create_at" gorm:"column:create_at"`
}
