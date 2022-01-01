package model

import "github.com/jinzhu/gorm"

type Coin struct {
	gorm.Model
	Name           string  `json:"name" binding:"required"`
	Amount         int     `json:"amount" binding:"required"`
	Price          float32 `json:"price" binding:"required"`
	TransactionFee int     `json:"transactionFee" binding:"required"`
	User           User    `gorm:"foreignkey:UserID"`
	UserID         uint
}

func (Coin) TableName() string {
	return "coins"
}
